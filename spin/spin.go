package spin

// RunConfig is a configuration struct for spinning up a particular service
type RunConfig struct {
	Image   string
	Tag     string
	Name    string
	Persist bool
	EnvOut  map[string]string
	EnvIn   map[string]string
}

// Spinner is an interface to be implemented by service that need to be spun up
type Spinner interface {
	Spin(c *RunConfig) error
}
