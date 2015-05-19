package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Fire takes a set of CLI arguments and returns an exit code (and maybe an
// error).
func fire(args []string) (int, error) {
	fset := flag.NewFlagSet("", flag.ExitOnError)

	// setup some command line flags
	var address, certfile, keyfile string
	fset.StringVar(&address, "address", ":3003", "the address to listen on")
	fset.StringVar(&keyfile, "keyfile", "host.key", "the keyfile to use")
	fset.StringVar(&certfile, "certfile", "site.crt", "the certfile to use")

	// attempt to parse the flags
	if err := fset.Parse(args); err != nil {
		return 1, err
	}

	// begin listening
	fmt.Println("Listening to HTTPS on address: " + address)

	// this blocks unless the listen fails (in which case we return a failing
	// exit code and the resulting error).
	return 1, http.ListenAndServeTLS(address, certfile, keyfile,
		// create an anonymous http handler
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/plain")

			if _, err := rw.Write([]byte("Hello from TLS!")); err != nil {
				log.Printf("couldn't write to response: %v", err)
			}
		}))
}

// run the actual program, handling errors as appropriate.
func main() {
	code, err := fire(os.Args)
	if err != nil {
		fmt.Errorf("Fatal error encountered: %s", err.Error())
	}

	os.Exit(code)
}
