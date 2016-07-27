// Author: Zhang Yifa
// Email: yzhang3@linkernetworks.com

package main

import (
	"log"
	"os"
)

func main() {
	//call console
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute command error: %v", err)
		os.Exit(-1)
	}
}
