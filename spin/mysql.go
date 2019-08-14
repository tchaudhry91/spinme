package spin

import "context"

// MySQL spins a MySQL container with the given settings. Nil config uses defaults
func MySQL(ctx context.Context, c *SpinConfig) (SpinOut, error) {
	if c == nil {
		c = &SpinConfig{}
	}
	if c.Image == "" {
		c.Image = "mysql"
	}
	if c.Tag == "" {
		c.Tag = "latest"
	}
	if c.Name == "" {
		c.Name = buildName("mysql")
	}
	if len(c.Env) == 0 {
		// default password used if none is provided
		c.Env = append(c.Env, "MYSQL_ROOT_PASSWORD=password")
		// default database used if none is provided
		c.Env = append(c.Env, "MYSQL_DATABASE=testdb")
	}
	c.ExposedPorts = []string{
		"5432",
	}
	return Generic(ctx, c)
}
