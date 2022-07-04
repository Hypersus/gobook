

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		fetch.go
//	Author:			Zeke
//	Date:			2022.04.21
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	defer fmt.Fprintf(os.Stderr, "defer function\n")
	for _, url := range os.Args[1:] {
		if ! strings.HasPrefix(url,"https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "%v\n",resp.Status)
	}
}