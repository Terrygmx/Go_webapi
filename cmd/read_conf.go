package main

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

type Redis_conf struct {
	Nodes    []string `json:"nodes"`
	Password string   `json:"password"`
}
type Api_conf struct {
	Port string `json:"webAPI_port"`
}
type Conf struct {
	Webapi *Api_conf   `json:"webAPI_config"`
	Redis  *Redis_conf `json:"redis_config"`
}

func ReadConf(fn string) (conf *Conf, error error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	byt := []byte{}
	for {

		byt, err = reader.ReadBytes('|')

		if err == io.EOF {
			break
		}
	}
	if err = json.Unmarshal(byt, &conf); err != nil {
		return nil, err
	}
	return conf, nil

}
