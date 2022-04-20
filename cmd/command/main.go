package main

import (
	"github.com/seniorescobar/design-patterns/cmd/command/button"
	"github.com/seniorescobar/design-patterns/cmd/command/command"
	"github.com/seniorescobar/design-patterns/cmd/command/device"
)

func main() {
	tv := &device.TV{}

	onButton := &button.Button{
		Command: &command.On{Device: tv},
	}

	offButton := &button.Button{
		Command: &command.Off{Device: tv},
	}

	onButton.Press()
	offButton.Press()
}
