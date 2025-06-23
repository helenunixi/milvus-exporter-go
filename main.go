package main

import (
    "log"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "milvus-exporter-go/metrics"
)

func main() {
    registry := prometheus.NewRegistry()
    collector := metrics.NewMilvusCollector("localhost:19530")
    registry.MustRegister(collector)

    http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
    log.Println("Milvus Exporter running on :9100/metrics")
    log.Fatal(http.ListenAndServe(":9100", nil))
}
