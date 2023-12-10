# api sever query k-v from Redis cluster
## file list
- read_conf.go
- redis_cluster_conn.go
- api_portal.go
- redis_test.go
- config.json
- runApi.sh

### config.json and read_conf.go
config.json: configurations on web api address and Redis cluster
read_conf.go: parse configuration file, returns a struct which wraps both api conf and redis cluster conf

### redis_cluster_conn.go && redis_test.go
redis_cluster_conn.go: connect to Redis cluster, and defines get key and set value functions
redis_test.go: test redis_cluster_conn with hardcoded cluster settings

### api_port.go
construct web api portal, which contains healthz and the Redis query portal("post"). healthz is a simple portal which returns a simple string to tell if the server is alive. The query portal parses request body and do the query then responds.

### runApi.sh
an example to start a web api server

### api
compiled binary file, which can run on centOS7

# how to start
1. compile with target archetecture and OS configuration
2. modify conf.json to your own configuration
3. start api server like:
```sh
    ./api config/conf.json &
```
4. query server, request body:
```json
 {"serial": "123xxxxyyyy"}
```
response should look like:
```json
{"serial": "123xxxxyyyy","label1":"aaaa","label2":"bbbb","label3":"cccc"}

```
## to start multiple api servers
simply make multiple conf files, like conf_1 - conf_10, give each one a different port(or on a different host), and run the script for each instance:
```sh
    ./api config/conf_1.json &
```

