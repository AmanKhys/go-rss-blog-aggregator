package main

import "net/http"

//the specific function signature the way that the go std library expects to define a http hander 

func handlerErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 400, "something went wrong")
}