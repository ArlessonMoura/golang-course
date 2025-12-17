package main

import "fmt"


func temperatureConversionKelvinToCelsius(KelvinTemp float64) float64 {
    celsiusTemp := KelvinTemp - 273
    return celsiusTemp
}

func main() {    
    CelsiusWaterBoilingPoint := temperatureConversionKelvinToCelsius(373.15)    
    
    fmt.Printf("O valor aproximado do ponto de ebulição em Celsius convertido de dados em Kelvin é %.2f°C\n", CelsiusWaterBoilingPoint)
}