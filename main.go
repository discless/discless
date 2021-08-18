package main

import (
	"github.com/discless/discless/api"
	"net/http"
)



func main()  {

	functionHandler := http.HandlerFunc(api.Apply)
	botHandler := http.HandlerFunc(api.NewBot)
	http.Handle("/function", functionHandler)
	http.Handle("/bot", botHandler)
	http.ListenAndServe("localhost:6969", nil)
}