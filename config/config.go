// Author:: Jonathan Hartman (<j@p4nt5.com>)
//
// Copyright (C) 2014, Jonathan Hartman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package config implements a simple struct and some functions for configuring
the application from a JSON config file. A proper configuration file should
contain:

    {
      "supermarket": {
	      "endpoint": "URL of Supermarket instance, e.g. https://supermarket.getchef.com"
      },
      "twitter": {
        "api_key": "Your Twitter API key, aka 'consumer key'",
        "api_secret": "Your Twitter API secret, aka 'consumer secret'",
        "access_token": "Your Twitter access token",
        "access_token_secret": "Your Twitter access token secret"
      }
    }
*/
package config

import (
	"encoding/json"
	"io"
	"os"
)

// SupermarketConfig implements a data structure for config items specific to
// the Supermarket instance being connected to.
type Supermarket struct {
	Endpoint string `json:"endpoint"`
}

// TwitterConfig imeplements a data structure for config items specific to
// Twitter.
type Twitter struct {
	APIKey            string `json:"api_key"`
	APISecret         string `json:"api_secret"`
	AccessToken       string `json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret"`
}

// Struct Config implements a data structure for an application configuration.
type Config struct {
	Supermarket Supermarket `json:"supermarket"`
	Twitter     Twitter     `json:"twitter"`
}

// New initializes and returns a new Config struct derived from the passed-in
// config file path.
func New(filename string) (c *Config, err error) {
	c = new(Config)
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	err = decodeJSON(f, c)
	return
}

// decodeJSON accepts an IO reader and a Config struct and populates that struct
// with the data read in.
func decodeJSON(r io.Reader, c *Config) (err error) {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}
