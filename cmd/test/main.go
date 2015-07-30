package main

import (
	"fmt"
	"github.com/tbruyelle/giphy"
)

func main() {
	g := giphy.NewClient("dc6zaTOxFJmzC")
	r, h, err := g.Search("poney")
	if err != nil {
		fmt.Println(err, h)
	}
	fmt.Printf("SUCCESS %+v", r)
}
