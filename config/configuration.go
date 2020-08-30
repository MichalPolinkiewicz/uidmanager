package config

import "time"

type Configuration struct {
	Port    string

	//supported bidders map
	Bidders map[string]*BidderDetails
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		Port: ":9001",
		Bidders: map[string]*BidderDetails{
			"adform":  {Name: "adform"},
			"rubicon": {Name: "rubicon"},
		},
	}
	return cfg
}

type BidderDetails struct {
	Id   int
	Name string
	TTL  time.Duration
}
