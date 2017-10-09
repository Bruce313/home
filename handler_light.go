package main

import (
	"net/http"
	"fmt"
	"log"
)

const LIGHT_URL = "http://"
type LightHandler struct {

}
const status = false
func (h *LightHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	log.Println("got request", req.URL.String())
	writer.Write([]byte(fmt.Sprintf(`{"status" : %v}`, status)))
}