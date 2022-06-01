package libshuttle

type optionConfig struct {
	ShowHostName string `mapstructure:"showHostname"`
  ShowUserName string `mapstructure:"showUserName"`
	Layout string `mapstructure:"layout"`
	ShowTime string `mapstructure:"showTime"`
}


type promptConfig struct {
	Icon string `mapstructure:"icon"`
  BackgroundColor string `mapstructure:"bgColor"`
}


type Config struct {
	Prompt  promptConfig `mapstructure:"prompt"`
	Options optionConfig   `mapstructure:"options"`
}

