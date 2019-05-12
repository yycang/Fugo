package main

import (
	"fmt"
)

type Any interface {}
type Car struct {
	Model string
	Manufacturer string
	BuildYear int
	Price int
}

type Cars []*Car

func main()  {
	ford := &Car{"fiesta", "ford", 2008, 123111}
	bmw := &Car{"xl450", "bmw", 2010, 1200000}
	merc := &Car{"d600", "mercedes", 2009, 300000}
	bmw2 := &Car{"x 800", "bmw", 1999, 2001000}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "bmw") && (car.BuildYear > 2001)
	})

	fmt.Println("allcars", allCars)
	fmt.Println("new bwms", allNewBMWs)

	manufactures := []string{"ford", "aston martin", "landrover", "bmw"}
	sortedAppender, sortedCars := MakeSortedAppender(manufactures)
	allCars.Process(sortedAppender)
	fmt.Println(sortedCars)
	bmwCount := len(sortedCars["bmw"])
	fmt.Println("we have:", bmwCount, "bmws")
}

func (cs Cars) Process(f func(car *Car))  {
	for _, c := range cs {
		f(c)
	}
}

func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)

	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func (cs Cars) Map(f func(car *Car) Any) []Any  {
	result := make([]Any, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

func MakeSortedAppender(manufacturers []string) (func(car *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)

	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["default"] = make([]*Car, 0)

	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["default"] = append(sortedCars["default"], c)
		}
	}
	return appender, sortedCars
}