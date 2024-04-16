package startup

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func SetupConfig() {
	var env string
	flag.StringVar(&env, "env", "dev", "which environment the service is running on")
	flag.Parse()

	viper.AddConfigPath("./config")
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}
	viper.AutomaticEnv()
}
