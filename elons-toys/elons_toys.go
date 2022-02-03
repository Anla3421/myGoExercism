package elon

import (
	"fmt"
	"strconv"
)

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery = c.battery - c.batteryDrain
		c.distance = c.distance + c.speed
	}
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	result := strconv.Itoa(c.distance)
	output := fmt.Sprintf("Driven %v meters", result)
	return output
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	result := strconv.Itoa(c.battery)
	output := fmt.Sprintf("Battery at %v%%", result)
	return output
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	for {
		c.Drive()
		if trackDistance <= c.distance {
			return true
		}
		if c.battery < c.batteryDrain {
			return false
		}
	}
}
