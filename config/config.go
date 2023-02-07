package config

import (
	"sync"
)

var PointURL = ""
var ConfigMu = &sync.Mutex{}
var Config *config

type config struct {

	ChannelEmotesRefreshMaxRange  int 	 	 `json:"CHANNEL_EMOTES_REFRESH_MAX_RANGE"`

	ParserIndexedDuration 		  int64   	 `json:"PARSER_INDEXED_DURATION"`
	
	ParserSkipMessages      	  []string 	 `json:"PARSER_SKIP_MESSAGES"`
	ParserSkipUsers			  []string	 `json:"PARSER_SKIP_USERS"`
	ParserSkipWords 	          []string	 `json:"PARSER_SKIP_WORDS"`

	MessageStoreTime        	  int64		 `json:"MESSAGE_STORE_TIME"`

}

func Get() config {
	ConfigMu.Lock(); defer ConfigMu.Unlock()
	cfg := *Config
	return cfg
}
