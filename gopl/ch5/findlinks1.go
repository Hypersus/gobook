

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		findlinks1.go
//	Author:			Zeke
//	Date:			2022.08.19
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package findlinks

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node)  []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c=c.NextSibling {
		links = visit(links,c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}




