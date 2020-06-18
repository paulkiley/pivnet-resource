// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v5"
)

type Creator struct {
	CreateStub        func() (pivnet.Release, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct{}
	createReturns     struct {
		result1 pivnet.Release
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Creator) Create() (pivnet.Release, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.recordInvocation("Create", []interface{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *Creator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Creator) CreateReturns(result1 pivnet.Release, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *Creator) CreateReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *Creator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Creator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
