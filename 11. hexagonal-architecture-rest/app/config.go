package app

import (
	"customer_api_hex_arch/domain"
	"customer_api_hex_arch/logger"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Database domain.DatabaseConfig
	Logger   logger.LoggerConfig
}

func LoadConfig(config *Config) {
	v := viper.New()
	v.AutomaticEnv()
	v.AllowEmptyEnv(true)
	// v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigName("config")
	// v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s\n", err)
	}

	// fix: https://github.com/spf13/viper/issues/798
	for _, key := range v.AllKeys() {
		// fmt.Println(key, v.Get(key))
		v.Set(key, v.Get(key))
	}

	v.Unmarshal(&config)
	// fmt.Println("Config:", *config)
	// fmt.Println("Config.Database.DB:", config.Database.DB)
	// fmt.Println("Config.Database.DRIVER:", config.Database.DRIVER)
	// fmt.Println("Config.Database.HOST:", config.Database.HOST)
	// fmt.Println("Config.Database.USER:", config.Database.USER)
}
