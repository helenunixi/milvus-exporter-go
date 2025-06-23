package metrics

import (
    "context"
    "log"

    "github.com/milvus-io/milvus-sdk-go/v2/client"
    "github.com/prometheus/client_golang/prometheus"
)

type MilvusCollector struct {
    milvus client.Client
    collectionsTotal *prometheus.Desc
}

func NewMilvusCollector(addr string) *MilvusCollector {
    return &MilvusCollector{
        collectionsTotal: prometheus.NewDesc(
            "milvus_collections_total",
            "Total number of collections in Milvus",
            nil, nil,
        ),
    }
}

func (c *MilvusCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- c.collectionsTotal
}

func (c *MilvusCollector) Collect(ch chan<- prometheus.Metric) {
    if c.milvus == nil {
        cli, err := client.NewGrpcClient(context.Background(), "localhost:19530")
        if err != nil {
            log.Println("Connection error:", err)
            return
        }
        c.milvus = cli
    }

    cols, err := c.milvus.ListCollections(context.Background())
    if err != nil {
        log.Println("Error listing collections:", err)
        return
    }

    ch <- prometheus.MustNewConstMetric(
        c.collectionsTotal,
        prometheus.GaugeValue,
        float64(len(cols)),
    )
}
