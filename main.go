package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flagMode := flag.String("mode", "client", "app mode")
	flag.Parse()

	switch *flagMode {
	case "server":

		go func() {
			if err := NewServerStart(); err != nil {
				log.Fatalf("couldn't run the server: %s", err)
			}
		}()
		if err := StartHTTPProxy(); err != nil {
			log.Fatalf("couldn't run the gateway: %s", err)
		}

	case "client":

		client, err := NewClient()
		if err != nil {
			log.Fatalf("couldn't run the client: %s", err)
		}
		fmt.Println(client.Version())

		if err := client.Watch("5"); err != nil {
			log.Fatalf("couldn't get the video stream: %s", err)
		}

	default:
		log.Fatalf("unknown mode %s", *flagMode)
	}

}
