package main

import "testing"

func TestParseConfig(t *testing.T) {
	confPath := "./config.toml"
	_, err := ParseConfig(confPath)
	if err != nil {
		t.Error(err)
		return
	}

}
