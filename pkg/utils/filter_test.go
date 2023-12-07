package utils

import (
	"testing"
)

type Config struct {
	A string
	B string
}

func TestFilter(t *testing.T) {
	confs := []Config{
		{
			A: "1",
			B: "1",
		},
		{
			A: "2",
			B: "2",
		},
		{
			A: "3",
			B: "3",
		},
		{
			A: "4",
			B: "4",
		},
		{
			A: "5",
			B: "1",
		},
		{
			A: "6",
			B: "1",
		},
		{
			A: "7",
		},
	}
	t.Run("过滤切片", func(t *testing.T) {
		got := Filter[[]Config, Config](confs, func(c Config) bool {
			switch c.B {
			case "1", "":
				return false
			default:
				return true
			}
		})
		t.Logf("got len = %d\n", len(got))
		t.Logf("got = %+v\n", got)
	})
}
