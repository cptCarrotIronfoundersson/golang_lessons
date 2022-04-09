package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("host or port are missed")
	}
	cli := NewTelnetClient(net.JoinHostPort(args[0], args[1]), *timeout, os.Stdin, os.Stdout)
	if err := cli.Connect(); err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect: %v", err))
	}
	defer cli.Close()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := cli.Send(); err != nil {
			fmt.Fprintf(os.Stderr, "Error during send")
		} else {
			fmt.Fprintf(os.Stderr, "...EOF")
		}
		cancel()
	}()

	go func() {
		if err := cli.Receive(); err != nil {
			fmt.Fprintf(os.Stderr, "Error during receive")
		} else {
			fmt.Fprintf(os.Stderr, "...Connection was closed by peer")
		}
		cancel()
	}()

	<-ctx.Done()
}
