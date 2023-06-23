/**
* Author: Xiangyu Wu
* Date: 2023-06-22
* From: hyperledger/fabric/common/metrics/metricsfakse/counter.go
*/

package metricsfakes

import (
	"sync"

	"github.com/geistwelt/quarkx/common/metrics"
)

type Counter struct {
	AddStub        func(float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 float64
	}

	WithStub        func(...string) metrics.Counter
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		arg1 []string
	}
	withReturns struct {
		result1 metrics.Counter
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Counter
	}

	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Counter) Add(arg float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct{arg1 float64}{arg1: arg})
	fake.recordInvocation("Add", []interface{}{arg})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(arg)
	}
}

// AddCallCount 返回调用 Add 方法的次数。
func (fake *Counter) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

// AddCalls 设置 Add 的 stub。
func (fake *Counter) AddCalls(stub func(float64)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

// AddArgsForCall 获得第 i 次调用 Add 方法传入的 float64 类型参数。
func (fake *Counter) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Counter) With(args ...string) metrics.Counter {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct{arg1 []string}{arg1: args})
	fake.recordInvocation("With", []interface{}{args})
	fake.withMutex.Unlock()

	if fake.WithStub != nil {
		return fake.WithStub(args...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.withReturns
	return fakeReturns.result1
}

// WithCallCount 返回调用 With 方法的次数。
func (fake *Counter) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

// WithCalls 设置 With 的 stub。
func (fake *Counter) WithCalls(stub func(...string) metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = stub
}

// WithArgsForCall 返回第 i 次调用 With 方法是传入的 []string 类型参数。
func (fake *Counter) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].arg1
}

// WithReturns 设置 With 方法的返回值。
// TODO 不明白的是，调用此方法为什么要让 WithStub 等于 nil。
func (fake *Counter) WithReturns(result metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	fake.withReturns = struct{result1 metrics.Counter}{result1: result}
}

// WithReturnsOnCall 设置 With 方法的第 i 个返回值。
// TODO 不明白此方法为什么要让 WithStub 等于 nil。
func (fake *Counter) WithReturnsOnCall(i int, result metrics.Counter) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct{result1 metrics.Counter})
	}
	fake.withReturnsOnCall[i] = struct{result1 metrics.Counter}{result1: result}
}

// Invocation 返回 {Add: [][]interface{}{}, With: [][]interface{}{}}，即过去调用 Add 方法和 With 方法的历
// 史记录，包括调用上述两个方法时传入的参数。
func (fake *Counter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Counter) recordInvocation(key string, args []interface{}) {
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
