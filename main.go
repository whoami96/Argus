package main

import (
	"errors"
	"log"
	"os"

	"github.com/whoami96/argus/collector"
	"github.com/whoami96/argus/config"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cfg, err = config.LoadConfig("config.yml")
		}
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
	}

	log.Printf("Loaded config for %d instances", len(cfg.Instances))

	for _, instance := range cfg.Instances {
		if instance.Type == "grafana" {
			log.Printf("Downloading alerts from %s (%s) ...", instance.Name, instance.URL)

		alerts, err := collector.FetchGrafanaAlerts(instance.URL, instance.Token)
		if err != nil {
			log.Printf("Error while fetching alerts from instance %s: %v", instance.Name, err)
			continue
		}

		log.Printf("Success! Found %d active alerts", len(alerts))
		for _, alert := range alerts {
			log.Printf(" -> [%s] %s (instance: %d): %s", alert.State, alert.Name, alert.Count, alert.Summary)
		}
	}
}
}
