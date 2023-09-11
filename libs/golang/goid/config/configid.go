package config

import "fmt"

type ID = string

func NewID(service string, source string) ID {
     return ID(fmt.Sprintf("config-%s-%s", service, source))
}
