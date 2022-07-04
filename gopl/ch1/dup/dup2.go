

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
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filePaths := os.Args[1:]
	if len(filePaths) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _,filePath := range filePaths {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n",err)
				continue
			}
			countLines(file, counts)
		}
	}
	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		counts[input.Text()]++
	}
}
