package entity

type ConfigInterface interface {
	SaveConfig(config *Config) error
	FindAll() ([]*Config, error)
	FindAllByService(service string) ([]*Config, error)
	FindOneById(id string) (*Config, error)
}
