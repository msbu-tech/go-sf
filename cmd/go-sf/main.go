package main

import (
	"fmt"
	"flag"
	
	"github.com/msbu-tech/wundercafe/cmd/version"
	"github.com/msbu-tech/go-sf/framework"
)

var (
	helpPtr     *bool = flag.Bool("help", false, "Help")
	versionPtr  *bool = flag.Bool("version", false, "Version Info")
	newPtr    *string = flag.String("new", "demo", "New App")
	tplPtr    *string = flag.String("template", "service", "Template Path")
	// outputPtr *string = flag.String("output", "./", "Output Path")
)

func main() {

	flag.Usage = usage
	flag.Parse()

	if *versionPtr == true {
		fmt.Println("go-sf", version.Version)
		fmt.Println("Copyright (c) MSBU-Tech, 2016")
	} else if *helpPtr == true {
		usage()
	} else {
		usage()
	}

	fmt.Print("new app: \n")
	err := framework.NewApp(*newPtr, *tplPtr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success!\n")
}

func usage() {
	fmt.Println("Usage")
}