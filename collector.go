package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type authCollector struct {
	authMetric      *prometheus.Desc
	attributeMetric *prometheus.Desc
	durationMetric  *prometheus.Desc
}

func newAuthCollector() *authCollector {
	return &authCollector{
		authMetric: prometheus.NewDesc("radius_authentication_success",
			"Indicates a successful authentication against the radius server",
			[]string{"username"},
			nil,
		),
		attributeMetric: prometheus.NewDesc("radius_authentication_response_attributes",
			"List of attributes sent in the reponse",
			[]string{"username", "attribute"},
			nil,
		),
		durationMetric: prometheus.NewDesc("radius_authentication_duration",
			"Amount of time to authetnicate against radius server",
			[]string{"username"},
			nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *authCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.authMetric
	ch <- collector.durationMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *authCollector) Collect(ch chan<- prometheus.Metric) {
	for _, e := range cfg.Accounts {

		start := time.Now()
		paket, success := tryRadiusAuth(e.Username, e.Password)
		duration := time.Since(start)

		labels := []string{e.Username}
		ch <- prometheus.MustNewConstMetric(collector.authMetric, prometheus.GaugeValue, boolToFloat64(success), labels...)
		ch <- prometheus.MustNewConstMetric(collector.durationMetric, prometheus.GaugeValue, float64(duration.Seconds()), labels...)

		if success {
			vlan := getVLANfromResponse(paket)
			labels = []string{e.Username, "Tunnel-Private-Group-Id"}
			ch <- prometheus.MustNewConstMetric(collector.attributeMetric, prometheus.GaugeValue, float64(vlan), labels...)
		}
	}

}
