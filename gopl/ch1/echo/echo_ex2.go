

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
)

func main() {
	for idx,arg := range os.Args[1:] {
		fmt.Printf("%v,%v\n",idx,arg)
	}
}
