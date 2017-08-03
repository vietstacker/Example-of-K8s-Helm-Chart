package main

import (
	"fmt"
	"net/http"
	"goji.io"
	"goji.io/pat"
	"encoding/json"
	"os"
	"net"
)

type Response struct {
	Type string
	Message string
}

var responses = []Response{
	Response{
		Type: "Acquaintance",
		Message: "Hi There",
	},

	Response{
		Type: "Stranger",
		Message: "Welcome !!!",
	},

}

func allGreetings(w http.ResponseWriter, r *http.Request){
	output, _ := json.Marshal(responses)
	fmt.Fprint(w, string(output))
}

func greetingByType(w http.ResponseWriter, r *http.Request){
	greeting_type := pat.Param(r, "type")
	for _, b := range responses {
		if b.Type == greeting_type {
			output, _ := json.Marshal(b)
			fmt.Fprint(w, string(output))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func logging(h http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Received request: %v\n", r.URL)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)

}


func getIp() string{
        hostname, _ := os.Hostname()
        addr, err := net.LookupHost(hostname)
        if err != nil {
                fmt.Println("Unknow host")
        }else {
                fmt.Println("IP Address: ",addr)
                fmt.Println(addr)

        }
        return addr[0]

}

func main() {
        mux := goji.NewMux()
        mux.HandleFunc(pat.Get("/greetings"), allGreetings)
        mux.HandleFunc(pat.Get("/greeting/:type"), greetingByType)
        mux.Use(logging)
        ip := getIp()
        listening_addr := string(ip) + ":8181"
        http.ListenAndServe(listening_addr, mux)
}
