package libs

import "github.com/spf13/viper"

// Create private data struct to hold config options.
type AppConfig struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
}

// Read the config file from the current directory and marshal
// into the conf config struct.
func LoadConfig(configPath string) (*AppConfig, error) {
	conf := &AppConfig{}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("server.yml")
	err := viper.ReadInConfig()

	if err != nil {
		return conf, err
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
