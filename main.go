package main

import(
  "github.com/spf13/viper"
  "os"
  "fmt"
  lib "shuttle/libshuttle"
)
var(
  conf lib.Config
)
func main() {
  
	viper.AutomaticEnv()
	viper.SetConfigName("shuttle")
	viper.SetConfigType("toml") 
    
	viper.AddConfigPath("$HOME/.config")
  viper.AddConfigPath("$HOME/.config/shuttle")
  
  if err := viper.ReadInConfig() ; err != nil {
    fmt.Println("Error reading config file: ",err)
    os.Exit(1)
  }
  
  if err := viper.Unmarshal(&conf) ; err != nil {
    fmt.Println("Error un-marashaling config: ",err)
    os.Exit(1)
  }
  
}
