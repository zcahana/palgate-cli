package main

import (
	"log"
	"os"

	palgate "github.com/zcahana/palgate-sdk"
)

func main() {
	config, err := palgate.InitConfig()
	if err != nil {
		log.Fatalf("Error parsing configuration: %v", err)
	}

	err = config.Validate()
	if err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	client := palgate.NewClient(config)

	logResp, err := client.Log()
	if err != nil {
		log.Fatalf("Error executing palgate command: %v", err)
	}

	if logResp.Status != palgate.ResponseStatusSuccess {
		log.Fatalf("Error executing palgate command: status=%s, error=%s, message=%s",
			logResp.Status, logResp.Error, logResp.Message)
	}

	err = Print(logResp.Records, os.Stdout)
	if err != nil {
		log.Fatalf("Error processing palgate log records: %v", err)
	}
}
