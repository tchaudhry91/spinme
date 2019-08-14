package spin

import (
	"context"
)

// SpinMongo spins a Mongo Container for the given settings
func SpinMongo(ctx context.Context, c *SpinConfig) (SpinOut, error) {
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
	return SpinGeneric(ctx, c)
}
