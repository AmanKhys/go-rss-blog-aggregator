package main

import "net/http"

//the specific function signature the way that the go std library expects to define a http hander 

func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w, 200, struct {}{})
}