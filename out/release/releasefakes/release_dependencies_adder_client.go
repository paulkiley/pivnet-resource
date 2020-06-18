// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v5"
)

type ReleaseDependenciesAdderClient struct {
	AddReleaseDependencyStub        func(productSlug string, releaseID int, dependentReleaseID int) error
	addReleaseDependencyMutex       sync.RWMutex
	addReleaseDependencyArgsForCall []struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}
	addReleaseDependencyReturns struct {
		result1 error
	}
	GetReleaseStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	getReleaseReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseDependenciesAdderClient) AddReleaseDependency(productSlug string, releaseID int, dependentReleaseID int) error {
	fake.addReleaseDependencyMutex.Lock()
	fake.addReleaseDependencyArgsForCall = append(fake.addReleaseDependencyArgsForCall, struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}{productSlug, releaseID, dependentReleaseID})
	fake.recordInvocation("AddReleaseDependency", []interface{}{productSlug, releaseID, dependentReleaseID})
	fake.addReleaseDependencyMutex.Unlock()
	if fake.AddReleaseDependencyStub != nil {
		return fake.AddReleaseDependencyStub(productSlug, releaseID, dependentReleaseID)
	} else {
		return fake.addReleaseDependencyReturns.result1
	}
}

func (fake *ReleaseDependenciesAdderClient) AddReleaseDependencyCallCount() int {
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	return len(fake.addReleaseDependencyArgsForCall)
}

func (fake *ReleaseDependenciesAdderClient) AddReleaseDependencyArgsForCall(i int) (string, int, int) {
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	return fake.addReleaseDependencyArgsForCall[i].productSlug, fake.addReleaseDependencyArgsForCall[i].releaseID, fake.addReleaseDependencyArgsForCall[i].dependentReleaseID
}

func (fake *ReleaseDependenciesAdderClient) AddReleaseDependencyReturns(result1 error) {
	fake.AddReleaseDependencyStub = nil
	fake.addReleaseDependencyReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseDependenciesAdderClient) GetRelease(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("GetRelease", []interface{}{productSlug, releaseVersion})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(productSlug, releaseVersion)
	} else {
		return fake.getReleaseReturns.result1, fake.getReleaseReturns.result2
	}
}

func (fake *ReleaseDependenciesAdderClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *ReleaseDependenciesAdderClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.getReleaseArgsForCall[i].productSlug, fake.getReleaseArgsForCall[i].releaseVersion
}

func (fake *ReleaseDependenciesAdderClient) GetReleaseReturns(result1 go_pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *ReleaseDependenciesAdderClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.invocations
}

func (fake *ReleaseDependenciesAdderClient) recordInvocation(key string, args []interface{}) {
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
