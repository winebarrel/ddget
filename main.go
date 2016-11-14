package main

import (
	"ddget"
	"fmt"
	"log"
)

func init() {
	log.SetFlags(0)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	ddgf, err := ddget.ParseFlag()

	if err != nil {
		panic(err)
	}

	ddg := ddget.New(ddgf.Profile, ddgf.Region)
	item, err := ddg.GetItem(ddgf.Table, ddgf.ValueAttrNmae, ddgf.Key)

	if err != nil {
		panic(err)
	}

	if ddgf.NoNewLine {
		fmt.Print(item)
	} else {
		fmt.Println(item)
	}
}
