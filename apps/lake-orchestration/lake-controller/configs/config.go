package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver                  string `mapstructure:"DB_DRIVER"`
	DBPort                    string `mapstructure:"DB_PORT"`
	DBHost                    string `mapstructure:"DB_HOST"`
	DBUser                    string `mapstructure:"DB_USER"`
	DBPassword                string `mapstructure:"DB_PASSWORD"`
	DBName                    string `mapstructure:"DB_NAME"`
	RabbitMQProtocol          string `mapstructure:"RABBITMQ_DEFAULT_PROTOCOL"`
	RabbitMQHost              string `mapstructure:"RABBITMQ_DEFAULT_HOST"`
	RabbitMQPort              string `mapstructure:"RABBITMQ_DEFAULT_PORT"`
	RabbitMQUser              string `mapstructure:"RABBITMQ_DEFAULT_USER"`
	RabbitMQPassword          string `mapstructure:"RABBITMQ_DEFAULT_PASS"`
	RabbitMQVhost             string `mapstructure:"RABBITMQ_DEFAULT_VHOST"`
	RabbitMQConsumerQueueName string `mapstructure:"RABBITMQ_CONSUMER_QUEUE_NAME"`
	RabbitMQConsumerName      string `mapstructure:"RABBITMQ_CONSUMER_NAME"`
	RabbitMQExchange          string `mapstructure:"RABBITMQ_DEFAULT_EXCHANGE"`
	RabbitMQExchangeType      string `mapstructure:"RABBITMQ_DEFAULT_EXCHANGE_TYPE"`
	RabbitMQDlxName           string `mapstructure:"RABBITMQ_DEFAULT_DLX_NAME"`
	WebServerPort             string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort            string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort         string `mapstructure:"GRAPHQL_SERVER_PORT"`
	JWTSecret                 string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpireIn               int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth                 *jwtauth.JWTAuth
}

type Config *conf

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("lake-gateway-config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}
