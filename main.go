package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

type account struct {
	Username string
	Password string
}

type config struct {
	Address  string
	Port     int
	Secret   string
	Timeout  int
	Accounts []account
}

var (
	configPath string
	cfg        config
)

func newConfig() config {
	c := config{}
	c.Address = "localhost"
	c.Port = 1812
	c.Timeout = 3
	return c
}

func readConfigFile(cfg *config) {
	f, err := os.Open(configPath)
	if err != nil {
		log.Printf("Couldn't open config file: %s", err)
		return
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Printf("Couldn't decode content of config file: %s", err)
		return
	}
}

func init() {
	flag.StringVar(&configPath, "config", "./config.yaml", "path to the config file")
}

func main() {

	flag.Parse()
	cfg = newConfig()
	readConfigFile(&cfg)

	authCollector := newAuthCollector()
	prometheus.MustRegister(authCollector)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
