

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		dup1.go
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
	"strings"
)

func main() {
	counts := make(map[string]int)
	filePaths := os.Args[1:]
	for _,filePath := range filePaths {
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n",err)
			continue
		}
		for _, line := range strings.Split(string(data),"\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

