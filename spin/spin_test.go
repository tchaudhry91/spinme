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
	}
	cases := []TestCase{
		{Name: "mongo-vanilla", Svc: "mongo"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			s := spin.SpinnerFunc(spin.SpinMongo)
			out, err := s.Spin(context.Background(), c.Config)
			if err != nil {
				t.Errorf("Failed to spin: %s: %s", c.Name, err.Error())
			}
			t.Log(out)
		})
	}
}
