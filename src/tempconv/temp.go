package main

import (
	"flag"
	"fmt"

	"../package/tempconv"
)

var tempC = flag.Float64("t", 0.0, "Температура по Цельсию")

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Print("Введите температура по цельсию: ")
		fmt.Scan(tempC)
	}

	boilingF := tempconv.CToF(tempconv.Celsius(*tempC))
	boilingK := tempconv.CToK(tempconv.Celsius(*tempC))

	fmt.Println(tempconv.Celsius(*tempC).String())
	fmt.Println(boilingF.String())
	fmt.Println(boilingK.String())
}
