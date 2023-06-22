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

func (fake *Counter) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
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
