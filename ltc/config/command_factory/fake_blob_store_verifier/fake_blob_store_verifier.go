// This file was generated by counterfeiter
package fake_blob_store_verifier

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/config/command_factory"
	"github.com/cloudfoundry-incubator/lattice/ltc/config/dav_blob_store"
)

type FakeBlobStoreVerifier struct {
	VerifyStub        func(config dav_blob_store.Config) (authorized bool, err error)
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		config dav_blob_store.Config
	}
	verifyReturns struct {
		result1 bool
		result2 error
	}
}

func (fake *FakeBlobStoreVerifier) Verify(config dav_blob_store.Config) (authorized bool, err error) {
	fake.verifyMutex.Lock()
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		config dav_blob_store.Config
	}{config})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(config)
	} else {
		return fake.verifyReturns.result1, fake.verifyReturns.result2
	}
}

func (fake *FakeBlobStoreVerifier) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *FakeBlobStoreVerifier) VerifyArgsForCall(i int) dav_blob_store.Config {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].config
}

func (fake *FakeBlobStoreVerifier) VerifyReturns(result1 bool, result2 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ command_factory.BlobStoreVerifier = new(FakeBlobStoreVerifier)
