package device

import "fmt"

type TV struct {
	isRunning bool
}

func (tv *TV) On() {
	tv.isRunning = true
	fmt.Println("Turning TV on...")
}

func (tv *TV) Off() {
	tv.isRunning = false
	fmt.Println("Turning TV off...")
}
