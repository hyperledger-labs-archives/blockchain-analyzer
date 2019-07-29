// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	ccprovider "github.com/hyperledger/fabric/core/common/ccprovider"
	peer "github.com/hyperledger/fabric/protos/peer"
)

type FileSystemSupport struct {
	CheckInstantiationPolicyStub        func(*peer.SignedProposal, string, []byte) error
	checkInstantiationPolicyMutex       sync.RWMutex
	checkInstantiationPolicyArgsForCall []struct {
		arg1 *peer.SignedProposal
		arg2 string
		arg3 []byte
	}
	checkInstantiationPolicyReturns struct {
		result1 error
	}
	checkInstantiationPolicyReturnsOnCall map[int]struct {
		result1 error
	}
	GetChaincodeFromLocalStorageStub        func(string, string) (ccprovider.CCPackage, error)
	getChaincodeFromLocalStorageMutex       sync.RWMutex
	getChaincodeFromLocalStorageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getChaincodeFromLocalStorageReturns struct {
		result1 ccprovider.CCPackage
		result2 error
	}
	getChaincodeFromLocalStorageReturnsOnCall map[int]struct {
		result1 ccprovider.CCPackage
		result2 error
	}
	GetChaincodesFromLocalStorageStub        func() (*peer.ChaincodeQueryResponse, error)
	getChaincodesFromLocalStorageMutex       sync.RWMutex
	getChaincodesFromLocalStorageArgsForCall []struct {
	}
	getChaincodesFromLocalStorageReturns struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	getChaincodesFromLocalStorageReturnsOnCall map[int]struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}
	GetInstantiationPolicyStub        func(string, ccprovider.CCPackage) ([]byte, error)
	getInstantiationPolicyMutex       sync.RWMutex
	getInstantiationPolicyArgsForCall []struct {
		arg1 string
		arg2 ccprovider.CCPackage
	}
	getInstantiationPolicyReturns struct {
		result1 []byte
		result2 error
	}
	getInstantiationPolicyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PutChaincodeToLocalStorageStub        func(ccprovider.CCPackage) error
	putChaincodeToLocalStorageMutex       sync.RWMutex
	putChaincodeToLocalStorageArgsForCall []struct {
		arg1 ccprovider.CCPackage
	}
	putChaincodeToLocalStorageReturns struct {
		result1 error
	}
	putChaincodeToLocalStorageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FileSystemSupport) CheckInstantiationPolicy(arg1 *peer.SignedProposal, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.checkInstantiationPolicyMutex.Lock()
	ret, specificReturn := fake.checkInstantiationPolicyReturnsOnCall[len(fake.checkInstantiationPolicyArgsForCall)]
	fake.checkInstantiationPolicyArgsForCall = append(fake.checkInstantiationPolicyArgsForCall, struct {
		arg1 *peer.SignedProposal
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("CheckInstantiationPolicy", []interface{}{arg1, arg2, arg3Copy})
	fake.checkInstantiationPolicyMutex.Unlock()
	if fake.CheckInstantiationPolicyStub != nil {
		return fake.CheckInstantiationPolicyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkInstantiationPolicyReturns
	return fakeReturns.result1
}

func (fake *FileSystemSupport) CheckInstantiationPolicyCallCount() int {
	fake.checkInstantiationPolicyMutex.RLock()
	defer fake.checkInstantiationPolicyMutex.RUnlock()
	return len(fake.checkInstantiationPolicyArgsForCall)
}

func (fake *FileSystemSupport) CheckInstantiationPolicyCalls(stub func(*peer.SignedProposal, string, []byte) error) {
	fake.checkInstantiationPolicyMutex.Lock()
	defer fake.checkInstantiationPolicyMutex.Unlock()
	fake.CheckInstantiationPolicyStub = stub
}

func (fake *FileSystemSupport) CheckInstantiationPolicyArgsForCall(i int) (*peer.SignedProposal, string, []byte) {
	fake.checkInstantiationPolicyMutex.RLock()
	defer fake.checkInstantiationPolicyMutex.RUnlock()
	argsForCall := fake.checkInstantiationPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FileSystemSupport) CheckInstantiationPolicyReturns(result1 error) {
	fake.checkInstantiationPolicyMutex.Lock()
	defer fake.checkInstantiationPolicyMutex.Unlock()
	fake.CheckInstantiationPolicyStub = nil
	fake.checkInstantiationPolicyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemSupport) CheckInstantiationPolicyReturnsOnCall(i int, result1 error) {
	fake.checkInstantiationPolicyMutex.Lock()
	defer fake.checkInstantiationPolicyMutex.Unlock()
	fake.CheckInstantiationPolicyStub = nil
	if fake.checkInstantiationPolicyReturnsOnCall == nil {
		fake.checkInstantiationPolicyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkInstantiationPolicyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorage(arg1 string, arg2 string) (ccprovider.CCPackage, error) {
	fake.getChaincodeFromLocalStorageMutex.Lock()
	ret, specificReturn := fake.getChaincodeFromLocalStorageReturnsOnCall[len(fake.getChaincodeFromLocalStorageArgsForCall)]
	fake.getChaincodeFromLocalStorageArgsForCall = append(fake.getChaincodeFromLocalStorageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetChaincodeFromLocalStorage", []interface{}{arg1, arg2})
	fake.getChaincodeFromLocalStorageMutex.Unlock()
	if fake.GetChaincodeFromLocalStorageStub != nil {
		return fake.GetChaincodeFromLocalStorageStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChaincodeFromLocalStorageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorageCallCount() int {
	fake.getChaincodeFromLocalStorageMutex.RLock()
	defer fake.getChaincodeFromLocalStorageMutex.RUnlock()
	return len(fake.getChaincodeFromLocalStorageArgsForCall)
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorageCalls(stub func(string, string) (ccprovider.CCPackage, error)) {
	fake.getChaincodeFromLocalStorageMutex.Lock()
	defer fake.getChaincodeFromLocalStorageMutex.Unlock()
	fake.GetChaincodeFromLocalStorageStub = stub
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorageArgsForCall(i int) (string, string) {
	fake.getChaincodeFromLocalStorageMutex.RLock()
	defer fake.getChaincodeFromLocalStorageMutex.RUnlock()
	argsForCall := fake.getChaincodeFromLocalStorageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorageReturns(result1 ccprovider.CCPackage, result2 error) {
	fake.getChaincodeFromLocalStorageMutex.Lock()
	defer fake.getChaincodeFromLocalStorageMutex.Unlock()
	fake.GetChaincodeFromLocalStorageStub = nil
	fake.getChaincodeFromLocalStorageReturns = struct {
		result1 ccprovider.CCPackage
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) GetChaincodeFromLocalStorageReturnsOnCall(i int, result1 ccprovider.CCPackage, result2 error) {
	fake.getChaincodeFromLocalStorageMutex.Lock()
	defer fake.getChaincodeFromLocalStorageMutex.Unlock()
	fake.GetChaincodeFromLocalStorageStub = nil
	if fake.getChaincodeFromLocalStorageReturnsOnCall == nil {
		fake.getChaincodeFromLocalStorageReturnsOnCall = make(map[int]struct {
			result1 ccprovider.CCPackage
			result2 error
		})
	}
	fake.getChaincodeFromLocalStorageReturnsOnCall[i] = struct {
		result1 ccprovider.CCPackage
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) GetChaincodesFromLocalStorage() (*peer.ChaincodeQueryResponse, error) {
	fake.getChaincodesFromLocalStorageMutex.Lock()
	ret, specificReturn := fake.getChaincodesFromLocalStorageReturnsOnCall[len(fake.getChaincodesFromLocalStorageArgsForCall)]
	fake.getChaincodesFromLocalStorageArgsForCall = append(fake.getChaincodesFromLocalStorageArgsForCall, struct {
	}{})
	fake.recordInvocation("GetChaincodesFromLocalStorage", []interface{}{})
	fake.getChaincodesFromLocalStorageMutex.Unlock()
	if fake.GetChaincodesFromLocalStorageStub != nil {
		return fake.GetChaincodesFromLocalStorageStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChaincodesFromLocalStorageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileSystemSupport) GetChaincodesFromLocalStorageCallCount() int {
	fake.getChaincodesFromLocalStorageMutex.RLock()
	defer fake.getChaincodesFromLocalStorageMutex.RUnlock()
	return len(fake.getChaincodesFromLocalStorageArgsForCall)
}

func (fake *FileSystemSupport) GetChaincodesFromLocalStorageCalls(stub func() (*peer.ChaincodeQueryResponse, error)) {
	fake.getChaincodesFromLocalStorageMutex.Lock()
	defer fake.getChaincodesFromLocalStorageMutex.Unlock()
	fake.GetChaincodesFromLocalStorageStub = stub
}

func (fake *FileSystemSupport) GetChaincodesFromLocalStorageReturns(result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.getChaincodesFromLocalStorageMutex.Lock()
	defer fake.getChaincodesFromLocalStorageMutex.Unlock()
	fake.GetChaincodesFromLocalStorageStub = nil
	fake.getChaincodesFromLocalStorageReturns = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) GetChaincodesFromLocalStorageReturnsOnCall(i int, result1 *peer.ChaincodeQueryResponse, result2 error) {
	fake.getChaincodesFromLocalStorageMutex.Lock()
	defer fake.getChaincodesFromLocalStorageMutex.Unlock()
	fake.GetChaincodesFromLocalStorageStub = nil
	if fake.getChaincodesFromLocalStorageReturnsOnCall == nil {
		fake.getChaincodesFromLocalStorageReturnsOnCall = make(map[int]struct {
			result1 *peer.ChaincodeQueryResponse
			result2 error
		})
	}
	fake.getChaincodesFromLocalStorageReturnsOnCall[i] = struct {
		result1 *peer.ChaincodeQueryResponse
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) GetInstantiationPolicy(arg1 string, arg2 ccprovider.CCPackage) ([]byte, error) {
	fake.getInstantiationPolicyMutex.Lock()
	ret, specificReturn := fake.getInstantiationPolicyReturnsOnCall[len(fake.getInstantiationPolicyArgsForCall)]
	fake.getInstantiationPolicyArgsForCall = append(fake.getInstantiationPolicyArgsForCall, struct {
		arg1 string
		arg2 ccprovider.CCPackage
	}{arg1, arg2})
	fake.recordInvocation("GetInstantiationPolicy", []interface{}{arg1, arg2})
	fake.getInstantiationPolicyMutex.Unlock()
	if fake.GetInstantiationPolicyStub != nil {
		return fake.GetInstantiationPolicyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getInstantiationPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileSystemSupport) GetInstantiationPolicyCallCount() int {
	fake.getInstantiationPolicyMutex.RLock()
	defer fake.getInstantiationPolicyMutex.RUnlock()
	return len(fake.getInstantiationPolicyArgsForCall)
}

func (fake *FileSystemSupport) GetInstantiationPolicyCalls(stub func(string, ccprovider.CCPackage) ([]byte, error)) {
	fake.getInstantiationPolicyMutex.Lock()
	defer fake.getInstantiationPolicyMutex.Unlock()
	fake.GetInstantiationPolicyStub = stub
}

func (fake *FileSystemSupport) GetInstantiationPolicyArgsForCall(i int) (string, ccprovider.CCPackage) {
	fake.getInstantiationPolicyMutex.RLock()
	defer fake.getInstantiationPolicyMutex.RUnlock()
	argsForCall := fake.getInstantiationPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FileSystemSupport) GetInstantiationPolicyReturns(result1 []byte, result2 error) {
	fake.getInstantiationPolicyMutex.Lock()
	defer fake.getInstantiationPolicyMutex.Unlock()
	fake.GetInstantiationPolicyStub = nil
	fake.getInstantiationPolicyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) GetInstantiationPolicyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getInstantiationPolicyMutex.Lock()
	defer fake.getInstantiationPolicyMutex.Unlock()
	fake.GetInstantiationPolicyStub = nil
	if fake.getInstantiationPolicyReturnsOnCall == nil {
		fake.getInstantiationPolicyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getInstantiationPolicyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorage(arg1 ccprovider.CCPackage) error {
	fake.putChaincodeToLocalStorageMutex.Lock()
	ret, specificReturn := fake.putChaincodeToLocalStorageReturnsOnCall[len(fake.putChaincodeToLocalStorageArgsForCall)]
	fake.putChaincodeToLocalStorageArgsForCall = append(fake.putChaincodeToLocalStorageArgsForCall, struct {
		arg1 ccprovider.CCPackage
	}{arg1})
	fake.recordInvocation("PutChaincodeToLocalStorage", []interface{}{arg1})
	fake.putChaincodeToLocalStorageMutex.Unlock()
	if fake.PutChaincodeToLocalStorageStub != nil {
		return fake.PutChaincodeToLocalStorageStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.putChaincodeToLocalStorageReturns
	return fakeReturns.result1
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorageCallCount() int {
	fake.putChaincodeToLocalStorageMutex.RLock()
	defer fake.putChaincodeToLocalStorageMutex.RUnlock()
	return len(fake.putChaincodeToLocalStorageArgsForCall)
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorageCalls(stub func(ccprovider.CCPackage) error) {
	fake.putChaincodeToLocalStorageMutex.Lock()
	defer fake.putChaincodeToLocalStorageMutex.Unlock()
	fake.PutChaincodeToLocalStorageStub = stub
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorageArgsForCall(i int) ccprovider.CCPackage {
	fake.putChaincodeToLocalStorageMutex.RLock()
	defer fake.putChaincodeToLocalStorageMutex.RUnlock()
	argsForCall := fake.putChaincodeToLocalStorageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorageReturns(result1 error) {
	fake.putChaincodeToLocalStorageMutex.Lock()
	defer fake.putChaincodeToLocalStorageMutex.Unlock()
	fake.PutChaincodeToLocalStorageStub = nil
	fake.putChaincodeToLocalStorageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemSupport) PutChaincodeToLocalStorageReturnsOnCall(i int, result1 error) {
	fake.putChaincodeToLocalStorageMutex.Lock()
	defer fake.putChaincodeToLocalStorageMutex.Unlock()
	fake.PutChaincodeToLocalStorageStub = nil
	if fake.putChaincodeToLocalStorageReturnsOnCall == nil {
		fake.putChaincodeToLocalStorageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putChaincodeToLocalStorageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FileSystemSupport) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkInstantiationPolicyMutex.RLock()
	defer fake.checkInstantiationPolicyMutex.RUnlock()
	fake.getChaincodeFromLocalStorageMutex.RLock()
	defer fake.getChaincodeFromLocalStorageMutex.RUnlock()
	fake.getChaincodesFromLocalStorageMutex.RLock()
	defer fake.getChaincodesFromLocalStorageMutex.RUnlock()
	fake.getInstantiationPolicyMutex.RLock()
	defer fake.getInstantiationPolicyMutex.RUnlock()
	fake.putChaincodeToLocalStorageMutex.RLock()
	defer fake.putChaincodeToLocalStorageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FileSystemSupport) recordInvocation(key string, args []interface{}) {
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