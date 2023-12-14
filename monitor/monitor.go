package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TxCounter         prometheus.Counter
	TokensTransferred prometheus.Gauge
)

func RegisterMetrics() {
	TxCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "usdt_transfer_transactions",
			Help: "Number of USDT transfer transactions.",
		},
	)
	TokensTransferred = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "usdt_tokens_transferred",
			Help: "Number of USDT tokens transferred.",
		},
	)
	prometheus.MustRegister(TxCounter, TokensTransferred)
}
