package request

import (
	"net/http"
)

type InterfaceRequest interface {
	Post(url string) (http.Response, error)
	Get(url string) (http.Response, error)
	Put(url string) (http.Response, error)
	Delete(url string) (http.Response, error)
}