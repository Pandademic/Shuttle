package main
import (
    "fmt"
    "runtime"
    "github.com/spf13/viper"
    "path/filepath"
     "strings"
     "os"
    "github.com/gookit/color"
)
func main() {
	var os string = runtime.GOOS
	viper.SetConfigName("shuttle") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")   // path to look for the config file in
	viper.AddConfigPath("$HOME")  
	viper.AddConfigPath("$HOME/.shuttle")
	viper.ReadInConfig() // Find and read the config file
	viper.SetDefault("prompt.icon", "$")
	viper.SetDefault("prompt.truncateDir",true)
	switch os{
		case "windows":
			var osLogo string= ""
			prompt(osLogo)
		case "darwin":
			var osLogo string= ""
			prompt(osLogo)
		case "linux":
			var osLogo string= ""
			prompt(osLogo)
	}
}
func trimPath(cwd, home string) string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		return path
	}
	pathSep := os.PathSeparator
	items := strings.Split(path,string(pathSep))
	truncItems := []string{}
	for i, item := range items {
		if i == (len(items) - 1) {
			truncItems = append(truncItems, item)
			break
		}
		truncItems = append(truncItems, item[:1])
	}
	return filepath.Join(truncItems...)
}
func prompt(osLogo string) {
	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render
	osSym := red(osLogo)
	if(osLogo == ""){
		osSym = cyan(osLogo)
	}else if(osLogo == ""){
		osSym = green(osLogo)
	}
	var prompt string = ""
	var icon string = viper.GetString("prompt.icon")
	//var yesTruncDir = viper.GetBool("prompt.truncateDir")
	prompt = "OS: "+osSym + "	"
	cwd , _ := os.Getwd()
	viper.AutomaticEnv()
	homeVar := viper.Get("HOME")
	prompt = prompt + red(""+trimPath(cwd,homeVar.(string))+" ")
	prompt = prompt + "" + cyan(icon) + "  "
	prompt = prompt + "" + cyan("◗") + ""
	fmt.Println(prompt)
}
