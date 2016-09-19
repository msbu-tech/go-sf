package main

import (
	"fmt"
	"flag"

	"github.com/msbu-tech/go-sf/cmd/version"
	"github.com/msbu-tech/go-sf/framework"
)

var (
	helpPtr     *bool = flag.Bool("help", false, "Help")
	versionPtr  *bool = flag.Bool("version", false, "Version Info")
	newPtr      *bool = flag.Bool("new", true, "New App")
	namePtr   *string = flag.String("name", "demo", "App Name")
	tplPtr    *string = flag.String("template", "service", "Template Path")
	outputPtr *string = flag.String("output", "./", "Output Path")
)

func main() {

	flag.Usage = usage
	flag.Parse()

	if *helpPtr == true {
		usage()

		return
	}
	if *versionPtr == true {
		fmt.Println("go-sf", version.Version)
		fmt.Println("Copyright (c) MSBU-Tech, 2016")

		return
	}

	if *newPtr == true {
		fmt.Println("New App")
		fmt.Println("\tApp Name: ", *namePtr)
		fmt.Println("\tTemplate Path: ", *tplPtr)
		fmt.Println("\tOutput Path: ", *outputPtr)

		err := framework.NewApp(*namePtr, *tplPtr)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Success.")
	}
}

func usage() {
	fmt.Println("Usage")
}