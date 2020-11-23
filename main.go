package main

import (
	"log"

	"github.com/pitakill/wttr/cmd"
)

func main() {
	o := &cmd.Options{
		A: true,
		F: true,
	}

	if err := cmd.NewClient(o).GetWeather(); err != nil {
		log.Fatalln(err)
	}
}
