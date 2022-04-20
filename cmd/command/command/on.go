package command

import "github.com/seniorescobar/design-patterns/cmd/command/device"

type On struct {
	Device device.Device
}

func (c *On) Execute() {
	c.Device.On()
}
