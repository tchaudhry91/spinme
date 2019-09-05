package spin

import (
	"context"
	"errors"
	"fmt"
	"strings"
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
	o, err := Generic(ctx, c)
	o.Service = "mongo"
	return o, err
}

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
		"3306",
	}
	o, err := Generic(ctx, c)
	o.Service = "mysql"
	return o, err
}

// MySQLConnString returns the sql open connection string from the spun container
func MySQLConnString(out SpinOut) (connStr string, err error) {
	var hostEp string
	var ok bool

	// Grab the host endpoint mapping for the container
	if hostEp, ok = out.Endpoints["3306/tcp"]; !ok {
		return "", errors.New("Failed to find proper port binding")
	}
	connStr = fmt.Sprintf("root:%s@tcp(%s)/%s", lookupEnv("MYSQL_ROOT_PASSWORD", out.Env), hostEp, lookupEnv("MYSQL_DATABASE", out.Env))
	return connStr, nil
}

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
	o, err := Generic(ctx, c)
	o.Service = "postgres"
	return o, err
}

// PostgresConnString returns the sql open connection string from the spun container
func PostgresConnString(out SpinOut) (connStr string, err error) {
	var hostEp string
	var ok bool

	// Grab the host endpoint mapping for the container
	if hostEp, ok = out.Endpoints["5432/tcp"]; !ok {
		return "", errors.New("Failed to find proper port binding")
	}
	// pq needs an independent port, not the entire endpoint
	ep := strings.Split(hostEp, ":")
	connStr = fmt.Sprintf("user=postgres password=%s dbname=%s port=%s sslmode=disable", lookupEnv("POSTGRES_PASSWORD", out.Env), lookupEnv("POSTGRES_DB", out.Env), ep[len(ep)-1])
	return connStr, nil
}

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
	o, err := Generic(ctx, c)
	o.Service = "redis"
	return o, err
}
