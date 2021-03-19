package models

import (
	"fmt"
	"github.com/BurntSushi/toml"

	log "github.com/sirupsen/logrus"
)

// CustomRequest is helper struct used for structuring fields in http request to api.etherscan.io
type CustomRequest struct {
	Tag    string
	APIKey string `toml:"API_KEY"`
}

// New method returns pointer to CustomRequest struct
//
//  `tag` passed in method as argument and encoded to hexdecimal formal
//
// API key automatically filling when this method called
func (c *CustomRequest) New(tag int) *CustomRequest {

	if tag <= 0 {
		log.Error("invalid string passed to method New(tag int) when try create CustomRequest")
	}
	c.setAPIKey()
	c.encodeTag(tag)
	return c
}

// encode `tag` argument to hex type with `0x` prefix
func(c *CustomRequest) encodeTag(tag int){
	hex := fmt.Sprintf("0x"+"%x",tag)
	c.Tag = hex
}
// parsed config.toml file and write API key to field APIKey
func (c *CustomRequest) setAPIKey() {
	if _, err := toml.DecodeFile("./config.toml", &c); err != nil {
		log.Error(err)
	}
}
