package iw

import (
	"github.com/prometheus/client_golang/prometheus"
)

type StationMetric struct {
	Desc      *prometheus.Desc
	ValueType prometheus.ValueType
	// TODO: could have a "expected second word", like {MBit/s,ms,seconds}
}

func newDesc(metric, desc string) *prometheus.Desc {
	return prometheus.NewDesc(prometheus.BuildFQName("iw", "", metric), desc, []string{"interface", "station"}, nil)
}

// StationMetrics maps raw field names to Prometheus metric desc.
var StationMetrics = map[string]StationMetric{
	"inactive time": {
		newDesc("inactive_time_ms", "milliseconds since last data exchange"),
		prometheus.CounterValue,
	},
	"rx bytes": {
		newDesc("receive_bytes_total", "bytes received"),
		prometheus.CounterValue,
	},
	"rx packets": {
		newDesc("receive_packets_total", "packets received"),
		prometheus.CounterValue,
	},
	"tx bytes": {
		newDesc("transmit_bytes_total", "bytes transmitted"),
		prometheus.CounterValue,
	},
	"tx packets": {
		newDesc("transmit_packets_total", "packets transmitted"),
		prometheus.CounterValue,
	},
	"tx retries": {
		newDesc("transmit_retries_total", "transmit retries"),
		prometheus.CounterValue,
	},
	"tx failed": {
		newDesc("transmit_failures_total", "transmit failures"),
		prometheus.CounterValue,
	},
	"signal": {
		newDesc("signal_strength_dbm", "signal strength in dBm"),
		prometheus.GaugeValue,
	},
	"signal avg": {
		newDesc("signal_strength_avg_dbm", "average signal strength in dBm"),
		prometheus.GaugeValue,
	},
	"tx bitrate": {
		newDesc("transmit_bitrate_mbps", "transmit bitrate in Mbps"),
		prometheus.GaugeValue,
	},
	"rx bitrate": {
		newDesc("receive_bitrate_mbps", "receive bitrate in Mbps"),
		prometheus.GaugeValue,
	},
	"authorized": {
		newDesc("is_authorized", "station is authorized (boolean)"),
		prometheus.GaugeValue,
	},
	"authenticated": {
		newDesc("is_authenticated", "station is authenticated (boolean)"),
		prometheus.GaugeValue,
	},
	"associated": {
		newDesc("is_associated", "station is associated (boolean)"),
		prometheus.GaugeValue,
	},
	"connected time": {
		newDesc("connected_seconds", "seconds that station has been connected"),
		prometheus.CounterValue,
	},
}
