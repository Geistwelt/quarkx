/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/metrics/statsd/provider.go
 */

package statsd

import (
	"github.com/geistwelt/quarkx/common/metrics"
	"github.com/geistwelt/quarkx/common/metrics/internal"
	"github.com/go-kit/kit/metrics/statsd"
)

const defaultFormat = "%{#fqname}"

type Provider struct {
	Statsd *statsd.Statsd
}

func (p *Provider) NewCounter(opt metrics.CounterOpts) metrics.Counter {
	if opt.StatsdFormat == "" {
		opt.StatsdFormat = defaultFormat
	}
	counter := &Counter{
		statsdProvider: p.Statsd,
		namer: internal.NewCounterNamer(opt),
	}

	if len(opt.LabelNames) == 0 {
		counter.Counter = p.Statsd.NewCounter(counter.namer.Format(), 1)
	}
	return counter
}

func (p *Provider) NewGauge(opt metrics.GaugeOpts) metrics.Gauge {
	if opt.StatsdFormat == "" {
		opt.StatsdFormat = defaultFormat
	}
	gauge := &Gauge{
		statsdProvider: p.Statsd,
		namer: internal.NewGaugeNamer(opt),
	}

	if len(opt.LabelNames) == 0 {
		gauge.Gauge = p.Statsd.NewGauge(gauge.namer.Format())
	}

	return gauge
}

func (p *Provider) NewHistogram(opt metrics.HistogramOpts) metrics.Histogram {
	if opt.StatsdFormat == "" {
		opt.StatsdFormat = defaultFormat
	}

	histogram := &Histogram{
		statsdProvider: p.Statsd,
		namer: internal.NewHistogramNamer(opt),
	}

	if len(opt.LabelNames) == 0 {
		histogram.Timing = p.Statsd.NewTiming(histogram.namer.Format(), 1)
	}

	return histogram
}

type Counter struct {
	Counter        *statsd.Counter
	namer          *internal.Namer
	statsdProvider *statsd.Statsd
}

func (c *Counter) Add(delta float64) {
	if c.Counter == nil {
		panic("label values must be provided by calling With")
	}
	c.Counter.Add(delta)
}

func (c *Counter) With(labelValues ...string) metrics.Counter {
	name := c.namer.Format(labelValues...)
	return &Counter{Counter: c.statsdProvider.NewCounter(name, 1)}
}

type Gauge struct {
	Gauge          *statsd.Gauge
	namer          *internal.Namer
	statsdProvider *statsd.Statsd
}

func (g *Gauge) Add(delta float64) {
	if g.Gauge == nil {
		panic("label values must be provided by calling With")
	}
	g.Gauge.Add(delta)
}

func (g *Gauge) Set(value float64) {
	if g.Gauge == nil {
		panic("label values must be provided by calling With")
	}
	g.Gauge.Set(value)
}

func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	name := g.namer.Format(labelValues...)
	return &Gauge{Gauge: g.statsdProvider.NewGauge(name)}
}

type Histogram struct {
	Timing         *statsd.Timing
	namer          *internal.Namer
	statsdProvider *statsd.Statsd
}

func (h *Histogram) Observe(value float64) {
	if h.Timing == nil {
		panic("label values must be provided by calling With")
	}
	h.Timing.Observe(value)
}

func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	name := h.namer.Format(labelValues...)
	return &Histogram{Timing: h.statsdProvider.NewTiming(name, 1)}
}