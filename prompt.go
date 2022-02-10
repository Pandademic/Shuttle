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
	viper.SetDefault("prompt.truncateDir",true)

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
	items := strings.Split(path, "/")
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
func winTrimPath(cwd, home string) string {
	var path string
	if strings.HasPrefix(cwd, home) {
		path = "~" + strings.TrimPrefix(cwd, home)
	} else {
		// If path doesn't contain $HOME, return the
		// entire path as is.
		path = cwd
		return path
	}
	items := strings.Split(path, "\\")
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
	osSym := red(osLogo)
	if osLogo == "" {
		osSym = cyan(osLogo)
	}else if osLogo == ""{
		osSym = green(osLogo)
	}
	var prompt string = ""
	var icon string = viper.GetString("prompt.icon")
	var yesTruncDir = viper.GetBool("prompt.truncateDir")
	prompt = "OS: "+osSym + " "
	if yesTrunc{
		switch os{
			case "windows":
				prompt = prompt + red(winTrimPath(os.Getwd,os.Getenv(HOME)))
			default:
				prompt = prompt + red(trimPath(os.Getwd,os.Getenv(HOME)))
		
		}
	}
	prompt = prompt + "" + cyan(icon)
	fmt.Println(prompt)
}
