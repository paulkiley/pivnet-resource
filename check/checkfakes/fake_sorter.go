// Code generated by counterfeiter. DO NOT EDIT.
package checkfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v5"
)

type FakeSorter struct {
	SortBySemverStub        func([]pivnet.Release) ([]pivnet.Release, error)
	sortBySemverMutex       sync.RWMutex
	sortBySemverArgsForCall []struct {
		arg1 []pivnet.Release
	}
	sortBySemverReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	sortBySemverReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	SortByLastUpdatedStub        func([]pivnet.Release) ([]pivnet.Release, error)
	sortByLastUpdatedMutex       sync.RWMutex
	sortByLastUpdatedArgsForCall []struct {
		arg1 []pivnet.Release
	}
	sortByLastUpdatedReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	sortByLastUpdatedReturnsOnCall map[int]struct {
		result1 []pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSorter) SortBySemver(arg1 []pivnet.Release) ([]pivnet.Release, error) {
	var arg1Copy []pivnet.Release
	if arg1 != nil {
		arg1Copy = make([]pivnet.Release, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.sortBySemverMutex.Lock()
	ret, specificReturn := fake.sortBySemverReturnsOnCall[len(fake.sortBySemverArgsForCall)]
	fake.sortBySemverArgsForCall = append(fake.sortBySemverArgsForCall, struct {
		arg1 []pivnet.Release
	}{arg1Copy})
	fake.recordInvocation("SortBySemver", []interface{}{arg1Copy})
	fake.sortBySemverMutex.Unlock()
	if fake.SortBySemverStub != nil {
		return fake.SortBySemverStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sortBySemverReturns.result1, fake.sortBySemverReturns.result2
}

func (fake *FakeSorter) SortBySemverCallCount() int {
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	return len(fake.sortBySemverArgsForCall)
}

func (fake *FakeSorter) SortBySemverArgsForCall(i int) []pivnet.Release {
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	return fake.sortBySemverArgsForCall[i].arg1
}

func (fake *FakeSorter) SortBySemverReturns(result1 []pivnet.Release, result2 error) {
	fake.SortBySemverStub = nil
	fake.sortBySemverReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeSorter) SortBySemverReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.SortBySemverStub = nil
	if fake.sortBySemverReturnsOnCall == nil {
		fake.sortBySemverReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.sortBySemverReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeSorter) SortByLastUpdated(arg1 []pivnet.Release) ([]pivnet.Release, error) {
	var arg1Copy []pivnet.Release
	if arg1 != nil {
		arg1Copy = make([]pivnet.Release, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.sortByLastUpdatedMutex.Lock()
	ret, specificReturn := fake.sortByLastUpdatedReturnsOnCall[len(fake.sortByLastUpdatedArgsForCall)]
	fake.sortByLastUpdatedArgsForCall = append(fake.sortByLastUpdatedArgsForCall, struct {
		arg1 []pivnet.Release
	}{arg1Copy})
	fake.recordInvocation("SortByLastUpdated", []interface{}{arg1Copy})
	fake.sortByLastUpdatedMutex.Unlock()
	if fake.SortByLastUpdatedStub != nil {
		return fake.SortByLastUpdatedStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.sortByLastUpdatedReturns.result1, fake.sortByLastUpdatedReturns.result2
}

func (fake *FakeSorter) SortByLastUpdatedCallCount() int {
	fake.sortByLastUpdatedMutex.RLock()
	defer fake.sortByLastUpdatedMutex.RUnlock()
	return len(fake.sortByLastUpdatedArgsForCall)
}

func (fake *FakeSorter) SortByLastUpdatedArgsForCall(i int) []pivnet.Release {
	fake.sortByLastUpdatedMutex.RLock()
	defer fake.sortByLastUpdatedMutex.RUnlock()
	return fake.sortByLastUpdatedArgsForCall[i].arg1
}

func (fake *FakeSorter) SortByLastUpdatedReturns(result1 []pivnet.Release, result2 error) {
	fake.SortByLastUpdatedStub = nil
	fake.sortByLastUpdatedReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeSorter) SortByLastUpdatedReturnsOnCall(i int, result1 []pivnet.Release, result2 error) {
	fake.SortByLastUpdatedStub = nil
	if fake.sortByLastUpdatedReturnsOnCall == nil {
		fake.sortByLastUpdatedReturnsOnCall = make(map[int]struct {
			result1 []pivnet.Release
			result2 error
		})
	}
	fake.sortByLastUpdatedReturnsOnCall[i] = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeSorter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sortBySemverMutex.RLock()
	defer fake.sortBySemverMutex.RUnlock()
	fake.sortByLastUpdatedMutex.RLock()
	defer fake.sortByLastUpdatedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSorter) recordInvocation(key string, args []interface{}) {
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
