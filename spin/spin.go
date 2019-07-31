package spin

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
	Expose       bool
	ExposedPorts []string
	Persist      bool
	PersistVols  []string
	EnvIn        []string
	EnvRemap     map[string]string
}

// SpinOut is an output structure containing values from the recently spun service
type SpinOut struct {
	ID     string
	IP     string
	Ports  nat.PortMap
	EnvOut map[string]string
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

// SpinGeneric is a generic spinner that assumes config input without modifying it
func SpinGeneric(ctx context.Context, c *SpinConfig) (SpinOut, error) {
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
	if c.Expose {
		var ports = make(nat.PortSet)
		for _, v := range c.ExposedPorts {
			p, err := nat.NewPort("tcp", v)
			if err != nil {
				return out, errors.Wrap(err, "Failed to parse expose port")
			}
			ports[p] = struct{}{}
		}
		cc.ExposedPorts = ports
	}
	if c.Persist {
		var vols = make(map[string]struct{})
		for _, v := range c.PersistVols {
			vols[v] = struct{}{}
		}
		cc.Volumes = vols
	}
	cc.Env = c.EnvIn
	ccb, err := cl.ContainerCreate(ctx, &cc, nil, nil, c.Name)
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
	out.IP = cInsp.NetworkSettings.IPAddress
	out.Ports = cInsp.NetworkSettings.Ports
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
