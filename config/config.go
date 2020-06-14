package config

import (
    "log"
    "os"

    "github.com/spf13/viper"
)

type Config struct {
    Debug       bool        `yaml:"debug"`
    Port        int         `yaml:"port"`
    Stage       string      `yaml:"stage"`
    LogConfig   LogConfig   `yaml:"logConfig"`
}

type LogConfig struct {
    LogLevel        string      `yaml:"logLevel"`
    LogTimeFormat   string      `yaml:"logtimeformat"`
}

func LoadConfig() *Config {
    var configFile, _ = os.LookupEnv("CONFIG")
    if(configFile == ""){
        log.Fatalln("No config file found, please set CONFIG variable correct")
    }
    log.Println("Loading configuration from ", configFile)
    
    viper.SetConfigType("yaml")
    viper.SetConfigFile(configFile)

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalln("Cannot read config file ", configFile)
    }

    config := &Config{}
    if err := viper.Unmarshal(config); err != nil {
        log.Fatalln("Cannot unmarshal config file ", err)
    }
    
    log.Print("Doing some iteration over the object ..")
    items := viper.AllSettings()
    for key, value := range items {
        log.Print("Key: ", key)
        log.Println(" Value ", value)
        //eachElementMap := value.(map[string]interface{})
        //for innerKey, innerVal := range eachElementMap {
        //    log.Print("innerKey: ", innerKey)
        //    log.Println(" innerVal: ", innerVal)
        //}
    }
    return config
}


