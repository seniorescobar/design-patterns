package button

import "github.com/seniorescobar/design-patterns/cmd/command/command"

type Button struct {
	Command command.Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
