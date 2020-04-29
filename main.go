package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr = flag.String("addr", "9090", "The addr of the application.")
	flag.Parse() // parse the flags

	c := newChannel()

	http.Handle("/", c)

	// get the channel/room going
	go c.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

