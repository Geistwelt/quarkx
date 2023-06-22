/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/metrics/prometheus/provider.go
 */

package prometheus

import (
	"github.com/geistwelt/quarkx/common/metrics"
	kitmetrics "github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	prom "github.com/prometheus/client_golang/prometheus"
)

type Provider struct{}

func (p *Provider) NewCounter(opt metrics.CounterOpts) metrics.Counter {
	return &Counter{
		Counter: prometheus.NewCounterFrom(
			prom.CounterOpts{
				Namespace: opt.Namespace,
				Subsystem: opt.Subsystem,
				Name: opt.Name,
				Help: opt.Help,
			},
			opt.LabelNames,
		),
	}
}

func (p *Provider) NewGauge(opt metrics.GaugeOpts) metrics.Gauge {
	return &Gauge{
		Gauge: prometheus.NewGaugeFrom(
			prom.GaugeOpts{
				Namespace: opt.Namespace,
				Subsystem: opt.Subsystem,
				Name: opt.Name,
				Help: opt.Help,
			},
			opt.LabelNames,
		),
	}
}

func (p *Provider) NewHistogram(opt metrics.HistogramOpts) metrics.Histogram {
	return &Histogram{
		Histogram: prometheus.NewHistogramFrom(
			prom.HistogramOpts{
				Namespace: opt.Namespace,
				Subsystem: opt.Subsystem,
				Name: opt.Name,
				Help: opt.Help,
				Buckets: opt.Buckets,
			},
			opt.LabelNames,
		),
	}
}

type Counter struct {
	kitmetrics.Counter
}

func (c *Counter) With(labelValues ...string) metrics.Counter {
	return &Counter{Counter: c.Counter.With(labelValues...)}
}

type Gauge struct {
	kitmetrics.Gauge
}

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	return &Gauge{Gauge: g.Gauge.With(labelValues...)}
}

type Histogram struct {
	kitmetrics.Histogram
}

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	return &Histogram{Histogram: h.Histogram.With(labelValues...)}
}