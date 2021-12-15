// # Aplicação mega simples somente para demostrar os conceitos de devOps
// # Constste em uma api qual irá consumir uma instancia REDIS para gerar respostas
// # Usando Golang devido a sua performance exemplar e simplicidade na construção de um WebServer simples

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func (_redis *RedisClient) connect_to_redis(redis_options *redis.Options) {
	_redis.client = redis.NewClient(redis_options)
}
func (_redis RedisClient) is_redis_connected(_context context.Context) bool {
	pong, err := _redis.client.Ping(_context).Result()
	fmt.Print("REDIS SAYS:")

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(pong)
		return true
	}
}

type Api struct {
	_redisClient RedisClient
}

func (_api *Api) getResponse(data map[string]string) []byte {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	return jsonResp
}

func (_api *Api) Helth_check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["status"] = "200"
	resp["uptime"] = uptime()
	resp["redis_connected"] = strconv.FormatBool(_api._redisClient.is_redis_connected(_api._redisClient.ctx))

	w.Write(_api.getResponse(resp))
}
func (_api *Api) insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	result, err := _api._redisClient.client.Set(_api._redisClient.ctx, key, value, 0).Result()
	if err != nil {
		resp["status"] = "500"
		resp["status_message"] = err.Error()
		http.Error(w, string(_api.getResponse(resp)), http.StatusInternalServerError)
	} else {
		resp["status"] = "200"
		resp["message"] = "KEY: " + key + " VALUE: " + value + " INSERIDA"
		fmt.Println(result)
		w.Write(_api.getResponse(resp))
	}
}

func (_api *Api) remove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	result, err := _api._redisClient.client.Del(_api._redisClient.ctx, key).Result()
	if err != nil {
		resp["status"] = "500"
		resp["status_message"] = err.Error()
		http.Error(w, string(_api.getResponse(resp)), http.StatusInternalServerError)
	} else {
		resp["status"] = "200"
		resp["message"] = "KEY: " + key + " VALUE: " + value + " REMOVIDA"
		fmt.Println(result)
		w.Write(_api.getResponse(resp))
	}
}

func (_api *Api) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	key := r.URL.Query().Get("key")

	result, err := _api._redisClient.client.Get(_api._redisClient.ctx, key).Result()
	if err != nil {
		resp["status"] = "404"
		resp["status_message"] = err.Error()
		http.Error(w, string(_api.getResponse(resp)), http.StatusNotFound)
	} else {
		resp["status"] = "200"
		resp["message"] = "KEY: " + key + " VALUE: " + _api._redisClient.client.Get(_api._redisClient.ctx, key).String()
		fmt.Println(result)
		w.Write(_api.getResponse(resp))
	}
}

var ctx = context.Background()
var startTime time.Time

func uptime() string {
	return time.Since(startTime).String()
}

func main() {
	// Main handler - basicamete nossa api de verdade...
	handler := http.NewServeMux()
	fmt.Println("INICIANDO SERVIDOR")

	// Aqui configuramos a conexão com o REDIS
	redisOptions := redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	_redis := RedisClient{ctx: ctx}
	_redis.connect_to_redis(&redisOptions)

	_api := Api{_redisClient: _redis}

	// Rotas basicas de chack
	handler.HandleFunc("/api/status", _api.Helth_check)

	// Rotas da aplicação CRUD
	handler.HandleFunc("/api/get", _api.get)
	handler.HandleFunc("/api/put", _api.insert)
	handler.HandleFunc("/api/delete", _api.remove)
	fmt.Println("INICIANDO SERVIDOR - NOW LISTENING")
	http.ListenAndServe("0.0.0.0:8080", handler)
	// Listen to all request to port 8080.
}
