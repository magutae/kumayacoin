package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	explorer "github.com/magutae/kumayacoin/explorer/templates"
	"github.com/magutae/kumayacoin/rest"
)

func usage() {
	fmt.Printf("Welcome to kumayacoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:		Start the Port of the server\n")
	fmt.Printf("-mode=rest:		Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {

	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
