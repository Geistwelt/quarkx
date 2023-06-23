/**
* Author: Xiangyu Wu
* Date: 2023-06-23
* From: hyperledger/fabric/common/metrics/metricsfakse/gauge.go
 */

package metricsfakes

import (
	"sync"

	"github.com/geistwelt/quarkx/common/metrics"
)

type Gauge struct {
	AddStub        func(float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg float64
	}

	SetStub        func(float64)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg float64
	}

	WithStub        func(...string) metrics.Gauge
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		arg []string
	}
	withReturns struct {
		result metrics.Gauge
	}
	withReturnsOnCall map[int]struct {
		result metrics.Gauge
	}

	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Gauge) Add(arg float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct{ arg float64 }{arg: arg})
	fake.recordInvocation("Add", []interface{}{arg})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(arg)
	}
}

// AddCallCount 返回过去调用 Add 方法的次数。
func (fake *Gauge) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

// AddCalls 为 Add 方法设置 stub。
func (fake *Gauge) AddCalls(stub func(float64)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

// AddArgsForCall 返回第 i 次调用 Add 方法时传入的参数。
func (fake *Gauge) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].arg
}

// Set Gauge 的 Set 方法。
func (fake *Gauge) Set(arg float64) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct{ arg float64 }{arg: arg})
	fake.recordInvocation("Set", []interface{}{arg})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(arg)
	}
}

// SetCallCount 返回过去调用 Set 方法的总次数。
func (fake *Gauge) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

// SetCalls 为 Set 方法设置 stub。
func (fake *Gauge) SetCalls(stub func(float64)) {
	fake.setMutex.Lock()
	defer fake.setMutex.Unlock()
	fake.SetStub = stub
}

// SetArgsForCall 返回第 i 次调用 Set 方法时传入的参数。
func (fake *Gauge) SetArgsForCall(i int) float64 {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].arg
}

// With Gauge 的 With 方法。
func (fake *Gauge) With(args ...string) metrics.Gauge {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct{ arg []string }{arg: args})
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

// WithCallCount 返回过去调用 With 方法的次数。
func (fake *Gauge) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

// WithCalls 为 With 方法设置 stub。
func (fake *Gauge) WithCalls(stub func(...string) metrics.Gauge) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = stub
}

// WithArgsForCall 返回第 i 次调用 With 方法时传入的参数。
func (fake *Gauge) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].arg
}

// WithReturns 直接设置 withReturns。
// TODO 不明白为什么要让 WithStub 等于 nil。
func (fake *Gauge) WithReturns(result metrics.Gauge) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	fake.withReturns = struct{ result metrics.Gauge }{result: result}
}

// WithReturnsOnCall 设置第 i 次调用 With 方法时需要返回的值。
func (fake *Gauge) WithReturnsOnCall(i int, result metrics.Gauge) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct{ result metrics.Gauge })
	}
	fake.withReturnsOnCall[i] = struct{ result metrics.Gauge }{result: result}
}

// Invocations 返回过去调用 Add Set 和 With 方法的历史记录，包括调用上述方法时传入的参数信息。
func (fake *Gauge) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Gauge) recordInvocation(key string, args []interface{}) {
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
