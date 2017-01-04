package main

import (
	"flag"
	"io"
	"net/http"

	"github.com/jamessanford/iw_exporter/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

var httpAddr = flag.String("http", ":6798", "listen on this address")

func main() {
	flag.Parse()

	prometheus.MustRegister(collector.NewIWCollector())

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, "iw_exporter\n")
	})
	http.Handle("/metrics", prometheus.Handler())
	log.Infof("listening on %v", *httpAddr)
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
