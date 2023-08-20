package entity

type SchemaInterface interface {
     SaveSchema(schema *Schema) error
     FindOneById(id string) (*Schema, error)
     FindAll() ([]*Schema, error)
     FindAllByService(service string) ([]*Schema, error)
}
