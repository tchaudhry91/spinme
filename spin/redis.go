package spin

import "context"

// Redis spins a Redis container with the given settings. Nil config uses defaults.
func Redis(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	if c == nil {
		c = &SpinConfig{}
	}
	if c.Image == "" {
		c.Image = "redis"
	}
	if c.Tag == "" {
		c.Tag = "latest"
	}
	if c.Name == "" {
		c.Name = buildName("redis")
	}
	c.ExposedPorts = []string{
		"6379",
	}
	return Generic(ctx, c)
}
