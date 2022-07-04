

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		echo4.go
//	Author:			Zeke
//	Date:			2022.04.25
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s","","separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}
