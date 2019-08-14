package spin_test

import (
	"context"
	"testing"

	"github.com/tchaudhry91/spinme/spin"
)

func TestSpinners(t *testing.T) {
	type TestCase struct {
		Name   string
		Svc    string
		Config *spin.SpinConfig
		Error  bool
	}
	cases := []TestCase{
		{Name: "mongo-vanilla", Svc: "mongo"},
		{Name: "mongo-with-name", Svc: "mongo", Config: &spin.SpinConfig{
			Name: "mongo-test-vanilla",
		}},
		{Name: "postgres-vanilla", Svc: "postgres"},
		{Name: "postgres-with-password", Svc: "postgres", Config: &spin.SpinConfig{
			Name: "postgres-custom",
			Env: []string{
				"POSTGRES_DB=newdb",
				"POSTGRES_PASSWORD=reallysecure",
			},
		}},
		{Name: "mysql-vanilla", Svc: "mysql"},
		{Name: "mysql-with-env", Svc: "mysql", Config: &spin.SpinConfig{
			Name: "mysql-custom",
			Env: []string{
				"MYSQL_ROOT_PASSWORD=pp22234222",
				"MYSQL_DATABASE=newdb",
			},
		}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			s, err := spin.SpinnerFrom(c.Svc)
			if err != nil {
				t.Errorf("Failed to find correct service")
				t.FailNow()
			}
			out, err := s.Spin(context.Background(), c.Config)
			if err != nil {
				t.Errorf("Failed to spin: %s: %s", c.Name, err.Error())
				t.FailNow()
			}
			t.Log(out)
			t.Log("Spun succesfully, slashing")
			err = spin.Slash(context.Background(), &out)
			if err != nil {
				t.Errorf("Failed to slash: %s: %s", out.ID, err.Error())
			}
		})
	}
}
