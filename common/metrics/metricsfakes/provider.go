/**
* Author: Xiangyu Wu
* Date: 2023-06-23
* From: hyperledger/fabric/common/metrics/metricsfakes/provider.go
 */

package metricsfakes

import (
	"sync"

	"github.com/geistwelt/quarkx/common/metrics"
)

type Provider struct {
	NewCounterStub        func(metrics.CounterOpts) metrics.Counter
	newCounterMutex       sync.RWMutex
	newCounterArgsForCall []struct {
		arg metrics.CounterOpts
	}
	newCounterReturns struct {
		result metrics.Counter
	}
	newCounterReturnsOnCall map[int]struct {
		result metrics.Counter
	}

	NewGaugeStub        func(metrics.GaugeOpts) metrics.Gauge
	newGaugeMutex       sync.RWMutex
	newGaugeArgsForCall []struct {
		arg metrics.GaugeOpts
	}
	newGaugeReturns struct {
		result metrics.Gauge
	}
	newGaugeReturnsOnCall map[int]struct {
		result metrics.Gauge
	}

	NewHistogramStub        func(metrics.HistogramOpts) metrics.Histogram
	newHistogramMutex       sync.RWMutex
	newHistogramArgsForCall []struct {
		arg metrics.HistogramOpts
	}
	newHistogramReturns struct {
		result metrics.Histogram
	}
	newHistogramReturnsOnCall map[int]struct {
		result metrics.Histogram
	}

	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Provider) NewCounter(arg metrics.CounterOpts) metrics.Counter {
	fake.newCounterMutex.Lock()
	ret, specificReturn := fake.newCounterReturnsOnCall[len(fake.newCounterArgsForCall)]
	fake.newCounterArgsForCall = append(fake.newCounterArgsForCall, struct{ arg metrics.CounterOpts }{arg: arg})
	fake.recordInvocation("NewCounter", []interface{}{arg})
	fake.newCounterMutex.Unlock()
	if fake.NewCounterStub != nil {
		return fake.NewCounterStub(arg)
	}
	if specificReturn {
		return ret.result
	}
	return fake.newCounterReturns.result
}

// NewCounterCallCount 返回过去调用 NewCounter 方法的次数。
func (fake *Provider) NewCounterCallCount() int {
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	return len(fake.newCounterArgsForCall)
}

// NewCounterCalls 为 NewCounter 方法设置 stub。
func (fake *Provider) NewCounterCalls(stub func(metrics.CounterOpts) metrics.Counter) {
	fake.newCounterMutex.Lock()
	defer fake.newCounterMutex.Unlock()
	fake.NewCounterStub = stub
}

// NewCounterArgsForCall 返回第 i 次调用 NewCounter 方法时传入的参数。
func (fake *Provider) NewCounterArgsForCall(i int) metrics.CounterOpts {
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	return fake.newCounterArgsForCall[i].arg
}

// NewCounterReturns 直接设置调用 NewCounter 方法时返回的值。
// TODO 为什么要让 NewCounterStub 等于 nil 呢？
func (fake *Provider) NewCounterReturns(result metrics.Counter) {
	fake.newCounterMutex.Lock()
	defer fake.newCounterMutex.Unlock()
	fake.NewCounterStub = nil
	fake.newCounterReturns = struct{ result metrics.Counter }{result: result}
}

// NewCounterReturnsOnCall 直接设置第 i 次调用 NewCounter 方法时会返回的值。
func (fake *Provider) NewCounterReturnsOnCall(i int, result metrics.Counter) {
	fake.newCounterMutex.Lock()
	defer fake.newCounterMutex.Unlock()
	fake.NewCounterStub = nil
	if fake.newCounterReturnsOnCall == nil {
		fake.newCounterReturnsOnCall = make(map[int]struct{ result metrics.Counter })
	}
	fake.newCounterReturnsOnCall[i] = struct{ result metrics.Counter }{result: result}
}

func (fake *Provider) NewGauge(arg metrics.GaugeOpts) metrics.Gauge {
	fake.newGaugeMutex.Lock()
	ret, specificReturn := fake.newGaugeReturnsOnCall[len(fake.newGaugeArgsForCall)]
	fake.recordInvocation("NewGauge", []interface{}{arg})
	fake.newGaugeMutex.Unlock()
	if fake.NewGaugeStub != nil {
		return fake.NewGaugeStub(arg)
	}
	if specificReturn {
		return ret.result
	}
	return fake.newGaugeReturns.result
}

// NewGaugeCallCount 返回过去调用 NewGauge 方法的总次数。
func (fake *Provider) NewGaugeCallCount() int {
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	return len(fake.newGaugeArgsForCall)
}

// NewGaugeCalls 为 NewGauge 方法设置 stub。
func (fake *Provider) NewGaugeCalls(stub func(metrics.GaugeOpts) metrics.Gauge) {
	fake.newGaugeMutex.Lock()
	defer fake.newGaugeMutex.Unlock()
	fake.NewGaugeStub = stub
}

// NewGaugeArgsForCall 返回第 i 次调用 NewGauge 方法时传入的参数。
func (fake *Provider) NewGaugeArgsForCall(i int) metrics.GaugeOpts {
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	return fake.newGaugeArgsForCall[i].arg
}

// NewGaugeReturns 直接设置调用 NewGauge 方法时返回的值。
func (fake *Provider) NewGaugeReturns(result metrics.Gauge) {
	fake.newGaugeMutex.Lock()
	defer fake.newGaugeMutex.Unlock()
	fake.NewGaugeStub = nil
	fake.newGaugeReturns = struct{ result metrics.Gauge }{result: result}
}

// NewGaugeReturnsOnCall 设置第 i 次调用 NewGuage 方法时会返回的值。
func (fake *Provider) NewGaugeReturnsOnCall(i int, result metrics.Gauge) {
	fake.newGaugeMutex.Lock()
	defer fake.newGaugeMutex.Unlock()
	fake.NewGaugeStub = nil
	if fake.newGaugeReturnsOnCall == nil {
		fake.newGaugeReturnsOnCall = make(map[int]struct{ result metrics.Gauge })
	}
	fake.newGaugeReturnsOnCall[i] = struct{ result metrics.Gauge }{result: result}
}

func (fake *Provider) NewHistogram(arg metrics.HistogramOpts) metrics.Histogram {
	fake.newHistogramMutex.Lock()
	ret, specificReturn := fake.newHistogramReturnsOnCall[len(fake.newCounterArgsForCall)]
	fake.newHistogramArgsForCall = append(fake.newHistogramArgsForCall, struct{arg metrics.HistogramOpts}{arg: arg})
	fake.recordInvocation("NewHistogram", []interface{}{arg})
	fake.newHistogramMutex.Unlock()
	if fake.NewHistogramStub != nil {
		return fake.NewHistogramStub(arg)
	}
	if specificReturn {
		return ret.result
	}
	return fake.newHistogramReturns.result
}

// NewHistogramCallCount 返回过去调用 NewHistogram 方法的次数。
func (fake *Provider) NewHistogramCallCount() int {
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	return len(fake.newHistogramArgsForCall)
}

// NewHistogramCalls 设置 NewHistogram 方法的 stub。
func (fake *Provider) NewHistogramCalls(stub func(metrics.HistogramOpts) metrics.Histogram) {
	fake.newHistogramMutex.Lock()
	defer fake.newHistogramMutex.Unlock()
	fake.NewHistogramStub = stub
}

// NewHistogramArgsForCall 返回第 i 次调用 NewHistogram 方法时传入的参数。
func (fake *Provider) NewHistogramArgsForCall(i int) metrics.HistogramOpts {
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	return fake.newHistogramArgsForCall[i].arg
}

// NewHistogramReturns 设置调用 NewHistogram 方法时返回的值。
func (fake *Provider) NewHistogramReturns(result metrics.Histogram) {
	fake.newHistogramMutex.Lock()
	defer fake.newHistogramMutex.Unlock()
	fake.NewHistogramStub = nil
	fake.newHistogramReturns = struct{result metrics.Histogram}{result: result}
}

// NewHistogramReturnsOnCall 设置第 i 次调用 NewHistogram 方法时会返回的值。
func (fake *Provider) NewHistogramReturnOnCall(i int, result metrics.Histogram) {
	fake.newHistogramMutex.Lock()
	defer fake.newHistogramMutex.Unlock()
	fake.NewHistogramStub = nil
	if fake.newHistogramReturnsOnCall == nil {
		fake.newHistogramReturnsOnCall = make(map[int]struct{result metrics.Histogram})
	}
	fake.newHistogramReturnsOnCall[i] = struct{result metrics.Histogram}{result: result}
}

func (fake *Provider) Invocations() map[string][][]interface{} {
	fake.newCounterMutex.RLock()
	defer fake.newCounterMutex.RUnlock()
	fake.newGaugeMutex.RLock()
	defer fake.newGaugeMutex.RUnlock()
	fake.newHistogramMutex.RLock()
	defer fake.newHistogramMutex.RUnlock()
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := make(map[string][][]interface{})
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Provider) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = make(map[string][][]interface{})
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = make([][]interface{}, 0)
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
