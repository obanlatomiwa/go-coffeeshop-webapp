package coffee

import (
	"fmt"
	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string
	Price float32
}

type CoffeeList struct {
	List []CoffeeDetails
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file, %w", err)
		return nil, err
	}

	err = viper.Unmarshal(&Coffees)
	if err != nil {
		return nil, err
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffeeName string) bool {
	for _, element := range Coffees.List {
		if element.Name == coffeeName {
			return true
		}
	}
	return false
}
