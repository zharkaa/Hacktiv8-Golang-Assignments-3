package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Wind  int
	Water int
}

func main() {
	rand.Seed(time.Now().Unix())

	for {
		time.Sleep(2 * time.Second)

		randomValueForWater := rand.Intn(100)
		randomValueForWind := rand.Intn(100)

		weather := Weather{Water: randomValueForWater, Wind: randomValueForWind}

		marsharledWeather, _ := json.MarshalIndent(weather, "", " ")

		cnvtWeather := string(marsharledWeather)

		fmt.Print("\n")
		fmt.Println(cnvtWeather)
		// fmt.Println("Water:", weather.Water, "Wind:", weather.Wind)

		if weather.Water <= 5 {
			fmt.Println("Status Water: Aman")
		} else if weather.Water >= 6 && weather.Water <= 8 {
			fmt.Println("Status Water: Siaga")
		} else if weather.Water > 8 {
			fmt.Println("Status Water: Bahaya")
		}

		if weather.Wind <= 6 {
			fmt.Println("Status Wind: Aman")
		} else if weather.Wind >= 7 && weather.Wind <= 15 {
			fmt.Println("Status Wind: Siaga")
		} else if weather.Wind > 15 {
			fmt.Println("Status Wind: Bahaya")
		}

	}

}
