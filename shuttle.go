package main

import(	
	c "github.com/pandadedemic/shuttle/pkg/config"
	pb "github.com/pandademic/shuttle/pkg/prompt_builder"
	"fmt"
)

func main(){
	config , err := c.CreateConf()
	if err != nil {
		panic(err)
	}
	prompt , err := pb.Build(config.EnabledSegments,config.Layout)
	if err != nil {
		panic(err)
	}
	fmt.Println(config.BeginIcon+prompt+config.EndIcon)
}
