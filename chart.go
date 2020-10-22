package flyt

import (
	"math"
)

type Chart struct {
	temps []int
}

func (c Chart) Temps() []int {
	return c.temps
}

func differenceAbs(temp1, temp2 int) int {
	return int(math.Abs(float64(temp1) - float64(temp2)))
}

func (c Chart) NearestTemp(temp int) int {
	var nearest = c.temps[0]
	var diff = differenceAbs(temp, c.temps[0])
	if len(c.temps) > 1 {
		for _, t := range c.temps[1:] {
			d := differenceAbs(temp, t)
			if d < diff {
				diff = d
				nearest = t
			}
		}
	}
	return nearest
}

func NewChart(temps ...int) Chart {
	return Chart{
		temps: temps,
	}
}
