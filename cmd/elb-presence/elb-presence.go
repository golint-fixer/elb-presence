// Copyright © 2014-2015, Civis Analytics

package main

import (
	"fmt"
	"github.com/civisanalytics/elb-presence/presence"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	var ELBName string

	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		switch kv[0] {
		case "ELB_NAME":
			ELBName = kv[1]
		}
	}

	if ELBName == "" {
		fmt.Printf("Environment Variable ELB_NAME must be set. Exiting.\n")
		os.Exit(1)
	}

	err := presence.JoinELB(ELBName)
	if err != nil {
		fmt.Printf("Unable to join Load-Balancer Pool[%s]\n", ELBName)
		os.Exit(1)
	}

	// Listen for system signals and adjust membership accordingly.
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, os.Kill)
	for {
		sc := <-signalChan
		switch sc {
		case syscall.SIGINT, syscall.SIGTERM, os.Kill:
			err := presence.LeaveELB(ELBName)
			if err != nil {
				fmt.Printf("Unable to leave Load-Balancer Pool[%s]\n", ELBName)
				os.Exit(1)
			}

			os.Exit(0)
		}
	}
}
