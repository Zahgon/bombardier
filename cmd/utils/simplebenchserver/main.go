package main

import (
	"bytes"
	"log"
	"net/http"

	"github.com/alecthomas/kingpin"
)

var serverPort = kingpin.Flag("port", "port to use for benchmarks").
	Default("8080").
	Short('p').
	String()
var responseSize = kingpin.Flag("size", "size of response in bytes").
	Default("1024").
	Short('s').
	Uint()

func main() {
	kingpin.Parse()
	response := bytes.Repeat([]byte("a"), int(*responseSize))
	addr := "localhost:" + *serverPort
	log.Println("Starting HTTP server on:", addr)
	lserr := http.ListenAndServe(addr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			_, werr := w.Write(response)
			if werr != nil {
				log.Println(werr)
			}
		}))
	if lserr != nil {
		log.Println(lserr)
	}
}
