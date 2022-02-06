package main
import (
    "fmt"
    "runtime"
     "github.com/spf13/viper"
)
func main() {
	type config struct {
		promptSymbol string
		Git bool
		multiLine bool
		topLine string // only if multi line = true
	}
	var os string = runtime.GOOS
	viper.SetConfigName("dotshuttle") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	err := viper.ReadInConfig() // Find and read the config file
	viper.SetDefault("promptSymbol", "$")
	viper.SetDefault("useGit",false)
	viper.SetDefault("multiLine",false)
	viper.SetDefault("topLine",false)// again only respected if multi-line = true
	var conf config 
	err := viper.Unmarshal(&C)
	if err != nil{
		fmt.Println("ERR: couldn't put push config into struct")
	}
	switch os {
				
		case "windows":
			windowsPrompt()
				
		case "darwin":
			darwinPrompt()
		case "linux":
			linuxPrompt()
	}

}
var (
	red   = color("\033[31m%s\033[0m")
	green = color("\033[32m%s\033[0m")
	cyan  = color("\033[36m%s\033[0m")
)

func color(s string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(s, fmt.Sprint(args...))
	}
}

var osLogo string = ""
func windowsPrompt(){
		osLogo = ""
		prompt()
}
func darwinPrompt(){
		osLogo = ""
		prompt()
}
func linuxPrompt() {
		osLogo = ""
		prompt()
}
func prompt() {
	fmt.Println(osLogo + " "+conf.promptSymbol)

}
