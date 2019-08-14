package spin

import "context"

// Postgres spins a Postgres container with the given settings. Nil config uses defaults
func Postgres(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	if c == nil {
		c = &SpinConfig{}
	}
	if c.Image == "" {
		c.Image = "postgres"
	}
	if c.Tag == "" {
		c.Tag = "alpine"
	}
	if c.Name == "" {
		c.Name = buildName("postgres")
	}
	if len(c.Env) == 0 {
		// default password used if none is provided
		c.Env = append(c.Env, "POSTGRES_PASSWORD=password")
		// default database used if none is provided
		c.Env = append(c.Env, "POSTGRES_DB=testdb")
	}
	c.ExposedPorts = []string{
		"5432",
	}
	return Generic(ctx, c)
}
