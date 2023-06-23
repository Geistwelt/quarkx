/**
* Author: Xiangyu Wu
* Date: 2023-06-23
* From: hyperledger/fabric/common/metrics/metricsfakes/histogram.go
 */

package metricsfakes

import (
	"sync"

	"github.com/geistwelt/quarkx/common/metrics"
)

type Histogram struct {
	ObserveStub        func(float64)
	observeMutex       sync.RWMutex
	observeArgsForCall []struct {
		arg float64
	}

	WithStub        func(...string) metrics.Histogram
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		arg []string
	}
	withReturns struct {
		result metrics.Histogram
	}
	withReturnsOnCall map[int]struct {
		result metrics.Histogram
	}

	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

// Observe Histogram 的 Observe 方法。
func (fake *Histogram) Observe(arg float64) {
	fake.observeMutex.Lock()
	fake.observeArgsForCall = append(fake.observeArgsForCall, struct{ arg float64 }{arg: arg})
	fake.recordInvocation("Observe", []interface{}{arg})
	fake.observeMutex.Unlock()
	if fake.ObserveStub != nil {
		fake.ObserveStub(arg)
	}
}

// ObserveCalls 为 Observe 方法设置 stub。
func (fake *Histogram) ObserveCalls(stub func(float64)) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = stub
}

// ObserveCallCount 返回过去调用 Observe 方法的次数。
func (fake *Histogram) ObserveCallCount() int {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return len(fake.observeArgsForCall)
}

// ObserveArgsForCall 返回第 i 次调用 Observe 方法时传入的参数。
func (fake *Histogram) ObserveArgsForCall(i int) float64 {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return fake.observeArgsForCall[i].arg
}

// With Histogram 的 With 方法，如果为 With 方法设置了 stub，则返回执行该 stub 产生的返回值；
// 如果为这次调用 With 方法设置了返回值，那么就返回设置的这个值。
func (fake *Histogram) With(args ...string) metrics.Histogram {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct{arg []string}{arg: args})
	fake.recordInvocation("With", []interface{}{args})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(args...)
	}
	if specificReturn {
		return ret.result
	}
	return fake.withReturns.result
}

// WithCallCount 返回过去调用 With 方法的总次数。
func (fake *Histogram) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

// WithCalls 为 With 方法设置 stub。
func (fake *Histogram) WithCalls(stub func(...string) metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = stub
}

// WithArgsForCall 返回第 i 次调用 With 方法时传入的参数。
func (fake *Histogram) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].arg
}

// WithReturns 直接设置调用 With 方法返回的值。
func (fake *Histogram) WithReturns(result metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	fake.withReturns = struct{result metrics.Histogram}{result: result}
}

// WithReturnsOnCall 设置第 i 次 With 方法的返回值。
func (fake *Histogram) WithReturnsOnCall(i int, result metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct{result metrics.Histogram})
	}
	fake.withReturnsOnCall[i] = struct{result metrics.Histogram}{result: result}
}

// Invocations 返回过去调用 Observe 和 With 方法的历史记录，包括当时调用前述
// 方法时传入的参数。
func (fake *Histogram) Invocations() map[string][][]interface{} {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := make(map[string][][]interface{})
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Histogram) recordInvocation(key string, args []interface{}) {
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
