package config

import(
	"github.com/mitchellh/go-homedir"
	"github.com/naoina/toml"
)

type ShuttleConfig struct{
	EnabledSegments []string
	Layout string
	BeginIcon string
	EndIcon string
}


func CreateConf() (ShuttleConfig,error){
	f , err := os.Open(homedir.Dir()+"/.config/shuttle.toml")
	if err != nil {
        	return nil,err
    	}
    	defer f.Close()
	var config ShuttleConfig
    	if err := toml.NewDecoder(f).Decode(&config); err != nil {
        	return nil,err
    	}

	return config , nil
}
