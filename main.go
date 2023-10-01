package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Weather struct {
	Water int `json:"Water"`
	Wind  int `json:"Wind"`
}

func (w *Weather) checkWeatherStatus() {
	if w.Water <= 5 {
		fmt.Println("Status Water: Aman")
	} else if w.Water >= 6 && w.Water <= 8 {
		fmt.Println("Status Water: Siaga")
	} else if w.Water > 8 {
		fmt.Println("Status Water: Bahaya")
	}

	if w.Wind <= 6 {
		fmt.Println("Status Wind: Aman")
	} else if w.Wind >= 7 && w.Wind <= 15 {
		fmt.Println("Status Wind: Siaga")
	} else if w.Wind > 15 {
		fmt.Println("Status Wind: Bahaya")
	}
}

func cnvtWeather(w *Weather) {
	res, _ := json.MarshalIndent(w, "", " ")
	fmt.Println(string(res))
}

func main() {
	weather := Weather{}

	// rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().Unix()))

	for {
		time.Sleep(2 * time.Second)

		weather.Water = 8
		weather.Wind = 15

		fmt.Println(weather)
		fmt.Print("\n")
		cnvtWeather(&weather)
		weather.checkWeatherStatus()

	}

}
