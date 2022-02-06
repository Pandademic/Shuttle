package main
import (
    "fmt"
    "runtime"
     "github.com/spf13/viper"
)
func main() {
	var os string = runtime.GOOS
	viper.SetConfigName("shuttle") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.ReadInConfig() // Find and read the config file
	viper.SetDefault("prompt.icon", "$")
	viper.SetDefault("git.enabled",true)
	switch os {			
		case "windows":
			var osLogoString string = ""
			prompt(osLogoString)
		case "darwin":
			var osLogoString string = ""
			prompt(osLogoString)
		case "linux":
			var osLogoString string = ""
			prompt(osLogoString)	
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
func prompt(osLogo string) {
	if osLogo == "" {
		var osSym string = cyan(osLogo)
	}else if osLogo == ""{
		var osSym string = red(osLogo)
	}else {
		var osSym string = green(osLogo)
	}
	fmt.Println(cyan(osLogo) + " "+viper.GetString("prompt.icon"))
}
