package main

//@todo: das muß alles noch schön gemacht werden

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	custumMetricsRegistry *prometheus.Registry
}

var (

	// counter example
	// Counter: Ein monotones Zählwerk, das nur erhöht werden kann. Es wird verwendet, um immer weiter wachsende Werte wie die Anzahl der Anfragen oder Fehler zu erfassen.
	MyCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_counter",
			Help: "A counter metric example",
		},
	)

	// gauge example
	// Gauge: Eine Metrik, die beliebig erhöht oder verringert werden kann. Sie wird verwendet, um aktuelle Zustände oder Werte wie Temperatur oder Speicherverbrauch zu erfassen.
	MyGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "my_gauge",
			Help: "A gauge metric example",
		},
	)

	// histogram example
	// Histogram: Eine Metrik, die Vorkommen und ihre Verteilung über konfigurierbare Buckets erfasst. Es wird oft verwendet, um Latenzen oder Größen von Anfragen zu messen.
	MyHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "my_histogram",
			Help:    "A histogram metric example",
			Buckets: prometheus.LinearBuckets(20, 5, 5), // Example: 5 Buckets, starting at 20, each 5 wide
		},
	)
	// summary example
	//Summary: Eine Metrik ähnlich wie Histogramm, die auch Verteilungen von Werten erfasst, aber perzentilenbasiert arbeitet.
	MySummary = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:       "my_summary",
			Help:       "A summary metric example",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}, // Example Objectives
		},
	)
)

func (m *metrics) registerMetrics() {
	// Register the metrics with Prometheus

	if err := m.custumMetricsRegistry.Register(MyCounter); err != nil {
		log.Fatalf("Failed to register custom metrics: %v", err)
	}
	if err := m.custumMetricsRegistry.Register(MyGauge); err != nil {
		log.Fatalf("Failed to register custom metrics: %v", err)
	}
	if err := m.custumMetricsRegistry.Register(MyHistogram); err != nil {
		log.Fatalf("Failed to register custom metrics: %v", err)
	}
	if err := m.custumMetricsRegistry.Register(MySummary); err != nil {
		log.Fatalf("Failed to register custom metrics: %v", err)
	}
}

func InitMetrics() {
	prometheus.MustRegister(MyCounter)
	prometheus.MustRegister(MyGauge)
	prometheus.MustRegister(MyHistogram)
	prometheus.MustRegister(MySummary)
}
