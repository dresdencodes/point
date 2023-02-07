package point 

import (
	"sync"
)

var purl = ""
var purlMu = &sync.Mutex{}

func pointsURL() string { purlMu.Lock(); defer purlMu.Unlock()

	if purl != "" {return purl}

	for key := range os.Args { current := os.Args[key]
		if strings.HasPrefix(current, "--point-url=")  {
			purl = strings.TrimPrefix(current, "--point-url=")
			break
		}
	} 
	if purl=="" {
		log.Fatal("--point-url opt must be provided")
	}
	if !strings.HasPrefix(purl, "http://") {purl = "http://"+url}
	if strings.HasSuffix(purl, "/") {
		log.Fatal("--point-url must not have trailing slash:", purl)
	}
	return purl
}