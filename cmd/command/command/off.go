package command

import "github.com/seniorescobar/design-patterns/cmd/command/device"

type Off struct {
	Device device.Device
}

func (c *Off) Execute() {
	c.Device.Off()
}
