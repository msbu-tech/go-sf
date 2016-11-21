package main

import (
	"flag"
	"fmt"

	"github.com/msbu-tech/go-sf/cmd/version"
	"github.com/msbu-tech/go-sf/framework"
)

var (
	helpPtr    *bool   = flag.Bool("help", false, "Help")
	versionPtr *bool   = flag.Bool("version", false, "Version Info")
	newPtr     *bool   = flag.Bool("new", true, "New App")
	namePtr    *string = flag.String("name", "demo", "App Name")
	authorPtr  *string = flag.String("author", "MSBU", "App Author")
	tplPtr     *string = flag.String("template", "service", "Template Path")
	outputPtr  *string = flag.String("output", "./", "Output Path")
)

func main() {

	flag.Usage = usage
	flag.Parse()

	if *helpPtr == true {
		usage()

		return
	} else if *versionPtr == true {
		fmt.Println("go-sf", version.Version)
		fmt.Println("Copyright (c) MSBU-Tech, 2016")

		return
	} else if *newPtr == true {
		fmt.Println("New App")
		fmt.Println("\tApp Name: ", *namePtr)
		fmt.Println("\tTemplate Path: ", *tplPtr)
		fmt.Println("\tOutput Path: ", *outputPtr)
		err, app := framework.New(*namePtr, *authorPtr, *tplPtr, *outputPtr)
		if err != nil {
			fmt.Println(err)
		}

		app.Generate()
	} else {
        usage()

        return
    }
}

func usage() {
	fmt.Println("go-sf Service Frame generator for Dpp in Go")
	fmt.Println("")
	fmt.Println("usage: go-sf [commands] [arguments]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("\t-new          # New app")
	fmt.Println("\t-help         # Show usage")
	fmt.Println("\t-version      # Show version")
	fmt.Println("")
	fmt.Println("Arguments:")
	fmt.Println("\t-name         # App name")
	fmt.Println("\t              # Default: demo")
	fmt.Println("\t-author       # Author name")
	fmt.Println("\t              # Default: MSBU")
	fmt.Println("\t-tpl          # Template path")
	fmt.Println("\t              # Default: service")
	fmt.Println("\t-output       # Output Path")
	fmt.Println("\t              # Default: ./")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("\tgo-sf -new -name=doctor -author=cscmaker -template=service -output=/home/dev/app/doctor")
	fmt.Println("")
}
