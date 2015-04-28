package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func fire(args []string) (int, error) {
	fset := flag.NewFlagSet("", flag.ExitOnError)

	var address, certfile, keyfile string
	fset.StringVar(&address, "address", ":3003", "the address to listen on")
	fset.StringVar(&keyfile, "keyfile", "rootCA.key", "the keyfile to use")
	fset.StringVar(&certfile, "certfile", "rootCA.crt", "the certfile to use")

	if err := fset.Parse(args); err != nil {
		return 1, err
	}

	fmt.Println("Listening to HTTPS on address: " + address)

	return 1, http.ListenAndServeTLS(address, certfile, keyfile, http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain")

		if _, err := rw.Write([]byte("Hello world!")); err != nil {
			log.Println("couldn't write to response: %v", err)
		}
	}))
}

func main() {
	code, err := fire(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error encountered: "+err.Error())
	}

	os.Exit(code)
}
