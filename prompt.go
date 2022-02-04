package main
import (
    "fmt"
    "runtime"
)
func main(){
		cwd := os.Getwd()
		home := os.Getenv("HOME")
		promptSym := "$"
		os := runtime.GOOS
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

var osLogo string = nil
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
	return fmt.Sprintf(
		"\n%s\n%s",
		promptSym,
	)
}
