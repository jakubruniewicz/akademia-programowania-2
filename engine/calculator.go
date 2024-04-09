package engine

import "math"

type Calculator struct {
	Sum  int
	Sub  int
	Mult int
	Div  float64
	Pow  float64
}

func (a *Calculator) Add(x, y int) {
	a.Sum = x + y
}

func (a *Calculator) Subtraction(x, y int) {
	a.Sub = x - y
}

func (a *Calculator) Multiplication(x, y int) {
	a.Mult = x * y
}

func (a *Calculator) Division(x, y float64) {
	a.Div = x / y
}
func (a *Calculator) Poww(x, y float64) {
	a.Pow = math.Pow(x, y)
}
