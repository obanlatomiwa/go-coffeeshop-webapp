package coffee

import (
	"testing"
)

func init() {
	Coffees = CoffeeList{
		List: []CoffeeDetails{
			CoffeeDetails{"Latte", 2.5},
			CoffeeDetails{"Flat White", 2},
			CoffeeDetails{"Cappucinno", 2.25},
		},
	}
}

func TestIsCoffeeAvailable(t *testing.T) {
	type testCase struct {
		coffeeType string
		available  bool
	}

	cases := []testCase{
		{"lat", false},
		{"Latte", true},
		{"", false},
		{"cappacinno", false},
	}

	for _, c := range cases {
		got := IsCoffeeAvailable(c.coffeeType)
		if c.available != got {
			t.Error("Expected ", c.available, "but got ", got)
		}
	}
}
