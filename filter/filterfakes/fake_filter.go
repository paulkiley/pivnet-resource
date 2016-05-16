// This file was generated by counterfeiter
package filterfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/pivnet-resource/filter"
	"github.com/pivotal-cf-experimental/pivnet-resource/pivnet"
)

type FakeFilter struct {
	DownloadLinksByGlobStub        func(downloadLinks map[string]string, glob []string) (map[string]string, error)
	downloadLinksByGlobMutex       sync.RWMutex
	downloadLinksByGlobArgsForCall []struct {
		downloadLinks map[string]string
		glob          []string
	}
	downloadLinksByGlobReturns struct {
		result1 map[string]string
		result2 error
	}
	DownloadLinksStub        func(p pivnet.ProductFiles) map[string]string
	downloadLinksMutex       sync.RWMutex
	downloadLinksArgsForCall []struct {
		p pivnet.ProductFiles
	}
	downloadLinksReturns struct {
		result1 map[string]string
	}
	invocations map[string][][]interface{}
}

func (fake *FakeFilter) DownloadLinksByGlob(downloadLinks map[string]string, glob []string) (map[string]string, error) {
	var globCopy []string
	if glob != nil {
		globCopy = make([]string, len(glob))
		copy(globCopy, glob)
	}
	fake.downloadLinksByGlobMutex.Lock()
	fake.downloadLinksByGlobArgsForCall = append(fake.downloadLinksByGlobArgsForCall, struct {
		downloadLinks map[string]string
		glob          []string
	}{downloadLinks, globCopy})
	fake.guard("DownloadLinksByGlob")
	fake.invocations["DownloadLinksByGlob"] = append(fake.invocations["DownloadLinksByGlob"], []interface{}{downloadLinks, globCopy})
	fake.downloadLinksByGlobMutex.Unlock()
	if fake.DownloadLinksByGlobStub != nil {
		return fake.DownloadLinksByGlobStub(downloadLinks, glob)
	} else {
		return fake.downloadLinksByGlobReturns.result1, fake.downloadLinksByGlobReturns.result2
	}
}

func (fake *FakeFilter) DownloadLinksByGlobCallCount() int {
	fake.downloadLinksByGlobMutex.RLock()
	defer fake.downloadLinksByGlobMutex.RUnlock()
	return len(fake.downloadLinksByGlobArgsForCall)
}

func (fake *FakeFilter) DownloadLinksByGlobArgsForCall(i int) (map[string]string, []string) {
	fake.downloadLinksByGlobMutex.RLock()
	defer fake.downloadLinksByGlobMutex.RUnlock()
	return fake.downloadLinksByGlobArgsForCall[i].downloadLinks, fake.downloadLinksByGlobArgsForCall[i].glob
}

func (fake *FakeFilter) DownloadLinksByGlobReturns(result1 map[string]string, result2 error) {
	fake.DownloadLinksByGlobStub = nil
	fake.downloadLinksByGlobReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) DownloadLinks(p pivnet.ProductFiles) map[string]string {
	fake.downloadLinksMutex.Lock()
	fake.downloadLinksArgsForCall = append(fake.downloadLinksArgsForCall, struct {
		p pivnet.ProductFiles
	}{p})
	fake.guard("DownloadLinks")
	fake.invocations["DownloadLinks"] = append(fake.invocations["DownloadLinks"], []interface{}{p})
	fake.downloadLinksMutex.Unlock()
	if fake.DownloadLinksStub != nil {
		return fake.DownloadLinksStub(p)
	} else {
		return fake.downloadLinksReturns.result1
	}
}

func (fake *FakeFilter) DownloadLinksCallCount() int {
	fake.downloadLinksMutex.RLock()
	defer fake.downloadLinksMutex.RUnlock()
	return len(fake.downloadLinksArgsForCall)
}

func (fake *FakeFilter) DownloadLinksArgsForCall(i int) pivnet.ProductFiles {
	fake.downloadLinksMutex.RLock()
	defer fake.downloadLinksMutex.RUnlock()
	return fake.downloadLinksArgsForCall[i].p
}

func (fake *FakeFilter) DownloadLinksReturns(result1 map[string]string) {
	fake.DownloadLinksStub = nil
	fake.downloadLinksReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeFilter) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ filter.Filter = new(FakeFilter)
