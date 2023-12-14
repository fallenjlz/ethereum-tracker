package main

import (
	"ethereum-tracker/handler"
	"ethereum-tracker/listener"
	"ethereum-tracker/monitor"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	EthereumNodeURL string `yaml:"ethereumNodeURL"`
	ContractAddress string `yaml:"contractAddress"`
}

func main() {
	var config Config
	configFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}

	listener.SetupListener(config.EthereumNodeURL, config.ContractAddress)
	monitor.RegisterMetrics()

	http.HandleFunc("/", handler.LatestLogsHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
