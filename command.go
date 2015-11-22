package potto

import (
	"fmt"
	"regexp"
	"strings"
)

type Command struct {
	Name string
	Args Args
}

func NewCommand(text, trigger string) *Command {
	pattern := fmt.Sprintf(`^(%s)`, trigger)

	text = regexp.MustCompile(pattern).ReplaceAllString(text, "")
	text = strings.TrimSpace(text)
	args := strings.Split(text, " ")

	return &Command{
		Name: args[0],
		Args: args[1:],
	}
}
