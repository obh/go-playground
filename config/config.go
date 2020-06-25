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
    LogConfig   LogConfig   `yaml:"logconfig"`
    DbConfig    DbConfig    `yaml:"dbconfig"`
}

type LogConfig struct {
    LogLevel        string      `yaml:"logLevel"`
    LogTimeFormat   string      `yaml:"logtimeformat"`
}

type DbConfig struct {
    Username        string      `yaml:"username"`
    Password        string      `yaml:"password"`
    Host            string      `yaml:"host"`
    Port            string      `yaml:"port"`
    Name            string      `yaml:"name"`
    MaxConnections  int         `yaml:"maxConnections"`
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
        //if key == "dbconfig" {
        //   LoadDbConfig(config, value.(map[string]interface{})) 
        //}
    }
    log.Println(config)
    return config
}

func LoadDbConfig(cfg *Config, val map[string]interface{}) {
    //dbConf := &DbConfig{}
    for k, v := range val {
        log.Print(k)
        log.Print(v)
    }
}


