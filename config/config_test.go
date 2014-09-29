package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var config_json = `
  {
    "supermarket": {
      "endpoint": "https://example.com"
    },
    "twitter": {
      "api_key": "abc",
      "api_secret": "123",
      "access_token": "doremi",
      "access_token_secret": "youandme"
    }
  }
`

func Test_New_1(t *testing.T) {
	f, err := ioutil.TempFile("", "supermarket-circular-config-test")
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}
	f.WriteString(config_json)
	defer f.Close()
	defer os.Remove(f.Name())

	c, err := New(f.Name())
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}
	for k, v := range map[string]string{
		c.Supermarket.Endpoint:      "https://example.com",
		c.Twitter.APIKey:            "abc",
		c.Twitter.APISecret:         "123",
		c.Twitter.AccessToken:       "doremi",
		c.Twitter.AccessTokenSecret: "youandme",
	} {
		if k != v {
			t.Fatalf("Expected: %s, got: %s", k, v)
		}
	}
}

func Test_decodeJSON_1(t *testing.T) {
	c := new(Config)
	r := strings.NewReader(config_json)

	err := decodeJSON(r, c)
	if err != nil {
		t.Fatalf("Expected no error, got: %s", err)
	}
	for k, v := range map[string]string{
		c.Supermarket.Endpoint:      "https://example.com",
		c.Twitter.APIKey:            "abc",
		c.Twitter.APISecret:         "123",
		c.Twitter.AccessToken:       "doremi",
		c.Twitter.AccessTokenSecret: "youandme",
	} {
		if k != v {
			t.Fatalf("Expected: %v, got: %v", v, k)
		}
	}
}
