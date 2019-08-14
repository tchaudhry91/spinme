package spin

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/phayes/freeport"

	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

// SpinConfig is a configuration struct for spinning up a particular service
type SpinConfig struct {
	Image        string
	Tag          string
	Name         string
	ExposedPorts []string
	Persist      bool
	PersistVols  []string
	Env          []string
}

// SpinOut is an output structure containing values from the recently spun service
type SpinOut struct {
	ID    string
	IP    string
	Ports nat.PortMap
	Env   []string
}

// Spinner is an interface to be implemented by service that need to be spun up
type Spinner interface {
	Spin(ctx context.Context, c *SpinConfig) (SpinOut, error)
}

// SpinnerFunc is a wrapper that allows using functions as Spinners
type SpinnerFunc func(ctx context.Context, c *SpinConfig) (SpinOut, error)

func (f SpinnerFunc) Spin(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	return f(ctx, c)
}

// Slash is a function to remove the given container
func Slash(ctx context.Context, o *SpinOut) error {
	cl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return errors.Wrap(err, "Failed to Create Docker Client from Environment")
	}
	err = cl.ContainerRemove(ctx, o.ID, types.ContainerRemoveOptions{Force: true})
	if err != nil {
		return errors.Wrap(err, "Failed to remove container")
	}
	return nil
}

// SpinGeneric is a generic spinner that assumes config input without modifying it
func Generic(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	var out SpinOut
	// Pull Image
	cl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return out, errors.Wrap(err, "Failed to Create Docker Client from Environment")
	}
	pull, err := cl.ImagePull(ctx, c.Image+":"+c.Tag, types.ImagePullOptions{})
	if err != nil {
		return out, errors.Wrap(err, "Failed to initiate pull image")
	}
	defer pull.Close()
	_, err = io.Copy(ioutil.Discard, pull)
	if err != nil {
		return out, errors.Wrap(err, "Failed to pull image")
	}
	cc := container.Config{
		Image: c.Image + ":" + c.Tag,
	}
	hc := container.HostConfig{}
	var ports = make(nat.PortSet)
	var portMap = make(nat.PortMap)
	for _, v := range c.ExposedPorts {
		p, err := nat.NewPort("tcp", v)
		if err != nil {
			return out, errors.Wrap(err, "Failed to parse expose port")
		}
		ports[p] = struct{}{}
		// Get a free port
		fp, err := freeport.GetFreePort()
		if err != nil {
			return out, errors.Wrap(err, "Failed to find a free port to bind to")
		}
		portMap[p] = []nat.PortBinding{
			nat.PortBinding{
				HostIP:   "0.0.0.0",
				HostPort: strconv.Itoa(fp),
			},
		}
	}
	hc.PortBindings = portMap
	cc.ExposedPorts = ports
	if c.Persist {
		var vols = make(map[string]struct{})
		for _, v := range c.PersistVols {
			vols[v] = struct{}{}
		}
		cc.Volumes = vols
	}
	cc.Env = c.Env
	ccb, err := cl.ContainerCreate(ctx, &cc, &hc, nil, c.Name)
	if err != nil {
		return out, errors.Wrap(err, "Error Creating Container")
	}
	if err = cl.ContainerStart(ctx, ccb.ID, types.ContainerStartOptions{}); err != nil {
		return out, errors.Wrap(err, "Error Starting Container")
	}
	cInsp, err := cl.ContainerInspect(ctx, ccb.ID)
	if err != nil {
		return out, errors.Wrap(err, "Error inspecting Started Container")
	}
	out.ID = ccb.ID
	out.IP = cInsp.NetworkSettings.IPAddress
	out.Ports = cInsp.NetworkSettings.Ports
	out.Env = cInsp.Config.Env
	return out, nil
}

// buildName returns a suitable name for the container
func buildName(svc string) string {
	var name string
	var dir string
	wd, err := os.Getwd()
	if err != nil {
		dir = "global"
	}
	if dir != "global" {
		dir, err = filepath.Abs(wd)
		if err != nil {
			dir = "global"
		}
		dir = filepath.Base(dir)
	}

	name = fmt.Sprintf("spinme-%s-%s-%d", dir, svc, time.Now().Unix())
	return name
}
