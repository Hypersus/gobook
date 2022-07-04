

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		server2.go
//	Author:			Zeke
//	Date:			2022.04.24
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"net/http"
	"fmt"
	"sync"
	"log"
	"example.com/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/counter",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

//func hanler(w http.ResponseWriter, r *http.Request) {
//	mu.Lock()
//	count++
//	mu.Unlock()
//	fmt.Fprintf(w, "URL.Path = %q\n",r.URL.Path)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

