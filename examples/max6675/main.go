package main

import (
	"machine"
	"max6675"
	"time"
)

func main() {
	thermo := max6675.New(machine.D5, machine.D6, machine.D7)
	thermo.Configure()

	for {

		if temp, err := thermo.ReadTemperature(); err != nil {
			println(err)
		} else {

			println(temp)
		}

		time.Sleep(time.Second * 1)
	}

}
