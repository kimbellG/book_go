package tempconv

import "fmt"

type Celsius float64
type Fathrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string     { return fmt.Sprintf("%g C", c) }
func (f Fathrenheit) String() string { return fmt.Sprintf("%g F", f) }
func (k Kelvin) String() string      { return fmt.Sprintf("%g K", k) }

//Функция преобразования температуры по цельсию к температуре по Фарренгейту
func CToF(c Celsius) Fathrenheit { return Fathrenheit(c*9/5 + 32) }

//Функция преобразования температуры по Фаренгейту к температуре по Цельсию
func FToC(f Fathrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//Преобразование Келвинов в Цельсии
func KToC(k Kelvin) Celsius { return Celsius(k - 273.1) }

//Преобразование Цельсия в Кельвины
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.1) }
