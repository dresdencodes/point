package config

import (
	"os"
	"log"
	"strings"

	"heiro/pkg/dg/dgreq"
	"heiro/pkg/dg/dgman"
)

var _ = dgman.InitRun("[every 5s] refresh config", func(){Load(false)})

func Load(firstLoad bool) {

	// point url empty
	if firstLoad {
		PointURL = findPointsURL() 
	}
	
	// locks
	ConfigMu.Lock(); defer ConfigMu.Unlock();
	
    err := dgreq.PostJSON(PointURL + "/points/config", nil, &Config)
    if err!=nil {
		// fatal on first load
		if firstLoad {
			log.Fatal("config request error ", err)
		}
        return 
    }
}

func findPointsURL() string {
	url := ""
	for key := range os.Args { current := os.Args[key]
		if strings.HasPrefix(current, "--point-url=")  {
			url = strings.TrimPrefix(current, "--point-url=")
			break
		}
	} 
	if url=="" {
		log.Fatal("--point-url opt must be provided")
	}
	if !strings.HasPrefix(url, "http://") {url = "http://"+url}
	if strings.HasSuffix(url, "/") {
		log.Fatal("--point-url must not have trailing slash:", url)
	}
	return url
}