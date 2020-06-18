// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v5"
)

type ReleaseFileGroupsAdder struct {
	AddReleaseFileGroupsStub        func(release pivnet.Release) error
	addReleaseFileGroupsMutex       sync.RWMutex
	addReleaseFileGroupsArgsForCall []struct {
		release pivnet.Release
	}
	addReleaseFileGroupsReturns struct {
		result1 error
	}
	addReleaseFileGroupsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseFileGroupsAdder) AddReleaseFileGroups(release pivnet.Release) error {
	fake.addReleaseFileGroupsMutex.Lock()
	ret, specificReturn := fake.addReleaseFileGroupsReturnsOnCall[len(fake.addReleaseFileGroupsArgsForCall)]
	fake.addReleaseFileGroupsArgsForCall = append(fake.addReleaseFileGroupsArgsForCall, struct {
		release pivnet.Release
	}{release})
	fake.recordInvocation("AddReleaseFileGroups", []interface{}{release})
	fake.addReleaseFileGroupsMutex.Unlock()
	if fake.AddReleaseFileGroupsStub != nil {
		return fake.AddReleaseFileGroupsStub(release)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addReleaseFileGroupsReturns.result1
}

func (fake *ReleaseFileGroupsAdder) AddReleaseFileGroupsCallCount() int {
	fake.addReleaseFileGroupsMutex.RLock()
	defer fake.addReleaseFileGroupsMutex.RUnlock()
	return len(fake.addReleaseFileGroupsArgsForCall)
}

func (fake *ReleaseFileGroupsAdder) AddReleaseFileGroupsArgsForCall(i int) pivnet.Release {
	fake.addReleaseFileGroupsMutex.RLock()
	defer fake.addReleaseFileGroupsMutex.RUnlock()
	return fake.addReleaseFileGroupsArgsForCall[i].release
}

func (fake *ReleaseFileGroupsAdder) AddReleaseFileGroupsReturns(result1 error) {
	fake.AddReleaseFileGroupsStub = nil
	fake.addReleaseFileGroupsReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseFileGroupsAdder) AddReleaseFileGroupsReturnsOnCall(i int, result1 error) {
	fake.AddReleaseFileGroupsStub = nil
	if fake.addReleaseFileGroupsReturnsOnCall == nil {
		fake.addReleaseFileGroupsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReleaseFileGroupsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseFileGroupsAdder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addReleaseFileGroupsMutex.RLock()
	defer fake.addReleaseFileGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseFileGroupsAdder) recordInvocation(key string, args []interface{}) {
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
