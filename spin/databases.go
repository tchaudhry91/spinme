/*
Copyright © 2019 Tanmay Chaudhry <tanmay.chaudhry@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	return Generic(ctx, c)
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
	return Generic(ctx, c)
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
	return Generic(ctx, c)
}
