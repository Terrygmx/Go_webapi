package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var conf *Conf
var err error
var rdb *redis.ClusterClient

type Serial struct {
	SerialNo string `json:"serial"`
}

// this handler is to detect if the api server is alive
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hello world!")
}

// parse request body and query from Redis
func post(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	r.Body.Read(body)
	var ser = &Serial{}
	json.Unmarshal(body, &ser)
	//	fmt.Printf("body:%s\n", body)

	var res string

	if res, err = GetStringFromCluster(rdb, ser.SerialNo); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, res)
}

func main() {
	// read configuration
	if conf, err = ReadConf(os.Args[1]); err != nil {
		panic(err)
	}
	//	fmt.Printf("conf read: %v, os.Arg: %s\n", conf, os.Args[1])
	apiPort := conf.Webapi.Port
	fmt.Printf("start web api server on port: %s\n", apiPort)
	nodes := conf.Redis.Nodes
	pwd := conf.Redis.Password
	// print only in debug mode, remember to delete this statement before deployment
	// fmt.Printf("nodes:%v, pwd: %s \n", nodes, pwd)

	clusterOptions := CreateClusterOptions(nodes, pwd)
	rdb = CreateCluster(clusterOptions)
	//	http.HandleFunc("/", healthz)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/post", post)
	http.ListenAndServe(fmt.Sprintf(":%s", apiPort), nil)
}
