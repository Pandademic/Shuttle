package main
import (
    "fmt"
    "runtime"
    //"os"
)
func main(){
		/*
		cwd,err:= os.Getwd()
		var home string = os.Getenv("HOME")
		var promptSym string = "$"
		*/
		var os string = runtime.GOOS
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

var osLogo string = "ﳑ"
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
	fmt.Println(osLogo)

}
