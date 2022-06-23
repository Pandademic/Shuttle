package prompt_builder

import (
	"strings"
	"errors"
	"os/user"
	"os"
	"time"
)

func buildUserSegment() string {
	return user.Username
}

func buildTimeSegment() string {
	dt := time.Now()
	return dt.Format("15:04:05")
}

func buildDirSegment() string {
	path, _ := os.Getwd()

	if path == user.HomeDir {
		return "~"
	}else{
	  return path
	}
}

func buildHostnameSegment() string {
	hostname , _ = os.Hostname
	return hostname
}

func Build(esegs []string,layout string) (string,error){
	lay := layout
	for _ , seg := range esegs {
		switch seg {
			case user:	
				lay = strings.replace(lay,"user",buildUserSegment(),-1)
			case time:
				lay = strings.replace(lay,"user",buildTimeSegment(),-1)
			case dir:
				lay = strings.replace(lay,"user",buildDirSegment(),-1)
			case hostname:
				lay = strings.replace(lay,"user",buildHostnameSegment(),-1)
			default:
				return nil,errors.New("Non regonized segment:"+seg)	
		}
	}
	return lay
}
