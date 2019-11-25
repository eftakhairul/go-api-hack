package libs

import "github.com/spf13/viper"

// AppConfig Struct
type AppConfig struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	DBName	string `yaml:"dbname"`
}

// LoadConfig reads the config file from the current directory and marshal
// into the conf config struct.
func LoadConfig(configPath string) (*AppConfig, error) {
	conf :=  &AppConfig{}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("server")
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
