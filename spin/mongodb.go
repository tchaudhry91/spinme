package spin

import (
	"context"
)

// Mongo spins a Mongo Container with the given settings. Nil config uses defaults
func Mongo(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	if c == nil {
		c = &SpinConfig{}
	}
	if c.Image == "" {
		c.Image = "mongo"
	}
	if c.Tag == "" {
		c.Tag = "latest"
	}
	if c.Name == "" {
		c.Name = buildName("mongo")
	}
	c.ExposedPorts = []string{
		"27017",
	}
	if len(c.Env) == 0 {
		// default user used if none is provided
		c.Env = append(c.Env, "MONGO_INITDB_ROOT_USERNAME=mongoadmin")
		// default password used if none is provided
		c.Env = append(c.Env, "MONGO_INITDB_ROOT_PASSWORD=password")
	}
	return Generic(ctx, c)
}
