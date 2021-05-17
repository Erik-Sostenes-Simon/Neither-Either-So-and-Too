package validate

import (
	"regexp"
	"strings"
)

type Message struct {
	Email   string
	Content string
	Errors  map[string]string
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(msg.Email))

	if match == false {
		msg.Errors["Email"] = "Por favor introduce un correo electronico valido"
	}
	if strings.TrimSpace(msg.Content) == "" {
		msg.Errors["Content"] = "Por favor introduce tu nombre"
	}
	return len(msg.Errors) == 0
}
