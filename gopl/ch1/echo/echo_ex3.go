

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		echo1.go
//	Author:			Zeke
//	Date:			2022.04.18
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start_slow := time.Now()
	var s,sep string
	for _,arg := range os.Args[1:] {
		s+=sep+arg
		sep=" "
	}
	end_slow := time.Now()
	elapsed_slow := end_slow.Sub(start_slow)
	fmt.Printf("time taken for slow version is %v \n",elapsed_slow)
	start_fast := time.Now()
	strings.Join(os.Args[1:]," ")
	end_fast := time.Now()
	elapsed_fast := end_fast.Sub(start_fast)
	fmt.Printf("time taken for fast version is %v \n",elapsed_fast)
}
