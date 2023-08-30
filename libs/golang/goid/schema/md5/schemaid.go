package schema

import "fmt"

type ID = string

func NewID(schemaType, service, source string) ID {
     return ID(fmt.Sprintf("%s-%s-%s", schemaType, service, source))
}
