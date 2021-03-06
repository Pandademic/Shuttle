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
var(
  osLogo string
  os string
)
func main() {
	//configure before anything
	viper.AutomaticEnv()
	viper.SetConfigName("shuttle")
	viper.SetConfigType("yaml") 
	// paths to look for the config file in
	viper.AddConfigPath(".")   
	viper.AddConfigPath("$HOME")  
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("$HOME/.config/shuttle")
	viper.ReadInConfig() // Find and read the config file
	viper.SetDefault("prompt.icon", "$")
	// detect enviorment
	os = runtime.GOOS
	switch os{
		case "windows":
			osLogo = ""
			prompt(osLogo)
		case "darwin":
			osLogo = ""
			prompt(osLogo)
		case "linux":
			osLogo = ""
			prompt(osLogo)
	}
}
func trimPath(cwd, home string) string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the entire path as is.
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
func use(vals ...interface{}){
    for _, val := range vals {
        _ = val
    }
}
func prompt(osLogo string) {
	// FG colors
	red := color.FgRed.Render
	green := color.FgGreen.Render
	cyan := color.FgCyan.Render
	white := color.FgWhite.Render
	blue := color.FgBlue.Render
	// BG colors
	bgYellow := color.BgYellow.Render
	bgRed := color.BgRed.Render
	bgGreen := color.BgGreen.Render
	bgCyan := color.BgCyan.Render
	bgWhite := color.BgWhite.Render
	bgBlue := color.BgBlue.Render
	use(red,blue,bgGreen,bgWhite,green) // don`t complain about unused colors
	osSym := bgRed(white(" "+osLogo+" "))
	if(osLogo == ""){
		osSym = bgBlue(white(" "+osLogo+" "))
	}else if(osLogo == ""){
		osSym = bgCyan(white(" "+osLogo+" "))
	}
	// render prompt
	var prompt string
	prompt = cyan("")
	if(viper.Get("prompt.segments.os" == true){
		prompt = prompt + osSym
	}
	if(viper.Get("prompt.segments.cwd") == true){
		cwd , _ := os.Getwd()
		homeVar := viper.Get("HOME")
		prompt = prompt + bgYellow(white(" :"+""+trimPath(cwd,homeVar.(string))+" "))\
	}
	// icon
	var icon string = viper.GetString("prompt.icon")
	prompt = prompt + "" + cyan(icon) + "  "
	// print it out
	fmt.Println(prompt)
}
