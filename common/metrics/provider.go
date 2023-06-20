/**
* Author: Xiangyu Wu
* Date: 2023-06-19
* From: hyperledger/fabric/common/metrics/provider.go
 */

package metrics

// Provider 是一个抽象接口，用来定义实例化 Counter、Gauge 和 Histogram 的方法。
type Provider interface {
	NewCounter(CounterOpts) Counter
	NewGauge(GaugeOpts) Gauge
	NewHistogram(HistogramOpts) Histogram
}

// Counter 是一个单调递增的计数器，它所记录的值只能增加，不能减少。例如，我们可以
// 用 Counter 表示访问 Http 服务的次数，不能利用 Counter 来统计电脑 CPU 的使用率。
type Counter interface {
	With(labelValues ...string) Counter
	Add(delta float64)
}

type CounterOpts struct {
	// Namespace, Subsystem, Name 被用来构成该 Counter 的全称。
	Namespace    string
	Subsystem    string
	Name         string

	// Help 提供该 Counter 的帮助信息。
	Help         string

	// TODO
	LabelNames   []string
	LabelHelp    map[string]string

	// TODO
	StatsdFormat string
}

// Gauge 用来度量既可增加又可减少的指标，例如 CPU 使用率和内存占用率等。
type Gauge interface {
	With(labelValues ...string) Gauge
	Add(delta float64)
	Set(value float64)
}

type GaugeOpts struct {
	Namespace string
	Subsystem string
	Name string

	Help string
	LabelNames []string

	LabelHelp map[string]string

	StatsdFormat string
}

// Histogram 可用来表示柱状图，用于统计数据的分布情况，例如统计一天中各个温
// 度区间的时间长等。
type Histogram interface {
	With(labelValues ...string) Histogram
	Observe(value float64)
}

type HistogramOpts struct {
	Namespace string
	Subsystem string
	Name string

	Help string
	Buckets []float64

	LabelNames []string
	LabelHelp map[string]string

	StatsdFormat string
}
