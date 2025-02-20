package router

import "net/http"

type Config struct {
	Addr string
}

type Router struct {
	Router *http.ServeMux
	config Config
}
