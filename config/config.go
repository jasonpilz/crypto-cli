package config

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

const app = "crypto"

var (
	CfgFile     string
	CfgFullPath string
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
}

// Init runs all setup steps for app configuration.
func Init() {
	SetEnv()
	SetConfig()
	SetDefaults()
	SetAliases()
	ReadConfig()
}

// SetConfig determines the paths viper will look for config files.
func SetConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		conf := fmt.Sprintf(".%v.yml", app)
		CfgFullPath = path.Join(home, conf)

		viper.AddConfigPath(home)
		viper.SetConfigName(conf)
	}
}

// ReadConfig loads the values from config file. If no file exists it will be created.
func ReadConfig() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.OpenFile(CfgFullPath, os.O_RDWR|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Created new config file: %s\n\n", CfgFullPath)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Using config file: %s\n\n", viper.ConfigFileUsed())
	}
}

// SetDefaults assigns default values for configuration values read from viper.
func SetDefaults() {
	// viper.SetDefault("size", 5)
}

// SetAliases creates a mechanism to access nested config vars by a shortened alias, to reduce density in the code.
func SetAliases() {
	// viper.RegisterAlias("amqp.c.gm", "amqp.consumer.goroutine_multiplier")
}

// SetEnv links the variables found in .env to be set/accessed by viper.
func SetEnv() {
	viper.SetEnvPrefix(app)

	// viper.BindEnv("es_url")

	viper.AutomaticEnv()
}
