package server

import (
	"fmt"
	"net/http"
	"sort"
)

type SimpleConn struct{}

func (sc *SimpleConn) AddRoutes() {
	http.HandleFunc("/url", sc.GetUrl)
	http.HandleFunc("/header", sc.GetHeader)
	http.HandleFunc("/method", sc.GetMethod)
}

func (sc *SimpleConn) GetUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.URL.String(): %s\n",r.URL.String())

	fmt.Fprintf(w, "r.URL.Path: %s\n", r.URL.Path)

	fmt.Fprintf(w, "r.URL.User.String(): %s\n", r.URL.User.String())
	fmt.Fprintf(w, "r.URL.User.Username(): %s\n", r.URL.User.Username())

	password, ok  := r.URL.User.Password()
	if ok {
		fmt.Fprintf(w, "r.URL.User.Password(): %s\n", password)
	} else {
		fmt.Fprintf(w, "r.URL.User.Password(): error\n")
	}

	fmt.Fprintf(w, "r.URL.Host: %s\n", r.URL.Host)

	fmt.Fprintf(w, "r.URL.Hostname: %s\n", r.URL.Hostname())
	fmt.Fprintf(w, "r.URL.Port: %s\n", r.URL.Port())
}

func (sc *SimpleConn) GetHeader(w http.ResponseWriter, r *http.Request) {
	headerKeys :=  make([]string, 0)
	for k := range r.Header {
		headerKeys =  append(headerKeys, k)
	}
	sort.Strings(headerKeys)

	for _, key := range headerKeys {
		fmt.Fprintf(w, "%s: %s\n", key, r.Header[key])
	}
}


func (sc *SimpleConn) GetMethod(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "r.Method = %s\n", r.Method)
}
