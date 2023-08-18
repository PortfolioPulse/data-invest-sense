package config

import "fmt"

type ID = string

func NewID(service string, source string) ID {
     return ID(fmt.Sprintf("%s-%s", service, source))
}
