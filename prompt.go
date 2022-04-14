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
// runtime variables
var(
  osLogo string = ""
  platform string = ""
  prompt string	 = ""
)
// config type
type config struct{
  icon string
  style string
  seperateSegments bool
  segmentSeperator string
  colorBasedOnExitCode bool
  showSomethingBeforePrompt bool
  somethingBeforePrompt string
}
// config instance
var(
   c config	
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
	// default options
	viper.SetDefault("icon", "$")
	viper.SetDefault("style", "plaintext")
	viper.SetDefault("seperateSegments", false)
	viper.SetDefault("segmentSeperator","")
	viper.SetDefault("colorBasedOnExitCode",false)
	viper.SetDefault("showSomethingBeforePrompt",false)
	viper.SetDefault("somethingBeforePrompt","")
	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("Unable to decode config")
	}
	// detect env
	platform = runtime.GOOS
	switch platform{
		case "windows":
			osLogo = ""
		case "darwin":
			osLogo = ""
		case "linux":
			osLogo = ""
	}
	deploy()
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
func deploy() {
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
	if(c.showSomethingBeforePrompt){
	    prompt = cyan(string(c.somethingBeforePrompt))
	}
	if(viper.Get("segments.os") == true){
		prompt = prompt + osSym
	}
	if(viper.Get("segments.cwd") == true){
		cwd , _ := os.Getwd()
		homeVar := os.Getenv("HOME")
		prompt = prompt + bgYellow(white(" :"+""+trimPath(cwd,string(homeVar)+" ")))
	}
	code := os.Getenv("?")
	if(c.colorBasedOnExitCode){
		if(platform == "windows"){
			if(code == "False"){
				prompt = prompt + "" + red(c.icon) + ""
			}else{
				prompt = prompt + "" + c.icon + "  "
			}
		}else{
			if(code != "0"){
				prompt = prompt + "" + red(c.icon) + "  "
			}else {
			    prompt = prompt + "" + c.icon + "  "
			}
		}
	}else{
		prompt = prompt + "" + c.icon + "  "
	}
	// print it out
	fmt.Println(prompt)
}
