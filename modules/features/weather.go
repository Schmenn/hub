package features

import (
	"fmt"
	"github.com/Schmenn/hub/structs"
)

// GetWeather returns weather data of the given coordinates or city
func GetWeather(i interface{}) {
	if s, ok := i.(string); ok {
		fmt.Println("string: ", s)
	} else if s, ok := i.(structs.Position); ok {
		fmt.Println("struct: ", s)
	} else {
		panic("")
	}
}
