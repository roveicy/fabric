// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/ledger"
)

type Lifecycle struct {
	ChaincodeDefinitionStub        func(chaincodeName string, txSim ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error)
	chaincodeDefinitionMutex       sync.RWMutex
	chaincodeDefinitionArgsForCall []struct {
		chaincodeName string
		txSim         ledger.QueryExecutor
	}
	chaincodeDefinitionReturns struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	chaincodeDefinitionReturnsOnCall map[int]struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}
	ChaincodeContainerInfoStub        func(chainID string, chaincodeID string) (*ccprovider.ChaincodeContainerInfo, error)
	chaincodeContainerInfoMutex       sync.RWMutex
	chaincodeContainerInfoArgsForCall []struct {
		chainID     string
		chaincodeID string
	}
	chaincodeContainerInfoReturns struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}
	chaincodeContainerInfoReturnsOnCall map[int]struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Lifecycle) ChaincodeDefinition(chaincodeName string, txSim ledger.QueryExecutor) (ccprovider.ChaincodeDefinition, error) {
	fake.chaincodeDefinitionMutex.Lock()
	ret, specificReturn := fake.chaincodeDefinitionReturnsOnCall[len(fake.chaincodeDefinitionArgsForCall)]
	fake.chaincodeDefinitionArgsForCall = append(fake.chaincodeDefinitionArgsForCall, struct {
		chaincodeName string
		txSim         ledger.QueryExecutor
	}{chaincodeName, txSim})
	fake.recordInvocation("ChaincodeDefinition", []interface{}{chaincodeName, txSim})
	fake.chaincodeDefinitionMutex.Unlock()
	if fake.ChaincodeDefinitionStub != nil {
		return fake.ChaincodeDefinitionStub(chaincodeName, txSim)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chaincodeDefinitionReturns.result1, fake.chaincodeDefinitionReturns.result2
}

func (fake *Lifecycle) ChaincodeDefinitionCallCount() int {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	return len(fake.chaincodeDefinitionArgsForCall)
}

func (fake *Lifecycle) ChaincodeDefinitionArgsForCall(i int) (string, ledger.QueryExecutor) {
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	return fake.chaincodeDefinitionArgsForCall[i].chaincodeName, fake.chaincodeDefinitionArgsForCall[i].txSim
}

func (fake *Lifecycle) ChaincodeDefinitionReturns(result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.ChaincodeDefinitionStub = nil
	fake.chaincodeDefinitionReturns = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeDefinitionReturnsOnCall(i int, result1 ccprovider.ChaincodeDefinition, result2 error) {
	fake.ChaincodeDefinitionStub = nil
	if fake.chaincodeDefinitionReturnsOnCall == nil {
		fake.chaincodeDefinitionReturnsOnCall = make(map[int]struct {
			result1 ccprovider.ChaincodeDefinition
			result2 error
		})
	}
	fake.chaincodeDefinitionReturnsOnCall[i] = struct {
		result1 ccprovider.ChaincodeDefinition
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeContainerInfo(chainID string, chaincodeID string) (*ccprovider.ChaincodeContainerInfo, error) {
	fake.chaincodeContainerInfoMutex.Lock()
	ret, specificReturn := fake.chaincodeContainerInfoReturnsOnCall[len(fake.chaincodeContainerInfoArgsForCall)]
	fake.chaincodeContainerInfoArgsForCall = append(fake.chaincodeContainerInfoArgsForCall, struct {
		chainID     string
		chaincodeID string
	}{chainID, chaincodeID})
	fake.recordInvocation("ChaincodeContainerInfo", []interface{}{chainID, chaincodeID})
	fake.chaincodeContainerInfoMutex.Unlock()
	if fake.ChaincodeContainerInfoStub != nil {
		return fake.ChaincodeContainerInfoStub(chainID, chaincodeID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.chaincodeContainerInfoReturns.result1, fake.chaincodeContainerInfoReturns.result2
}

func (fake *Lifecycle) ChaincodeContainerInfoCallCount() int {
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	return len(fake.chaincodeContainerInfoArgsForCall)
}

func (fake *Lifecycle) ChaincodeContainerInfoArgsForCall(i int) (string, string) {
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	return fake.chaincodeContainerInfoArgsForCall[i].chainID, fake.chaincodeContainerInfoArgsForCall[i].chaincodeID
}

func (fake *Lifecycle) ChaincodeContainerInfoReturns(result1 *ccprovider.ChaincodeContainerInfo, result2 error) {
	fake.ChaincodeContainerInfoStub = nil
	fake.chaincodeContainerInfoReturns = struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) ChaincodeContainerInfoReturnsOnCall(i int, result1 *ccprovider.ChaincodeContainerInfo, result2 error) {
	fake.ChaincodeContainerInfoStub = nil
	if fake.chaincodeContainerInfoReturnsOnCall == nil {
		fake.chaincodeContainerInfoReturnsOnCall = make(map[int]struct {
			result1 *ccprovider.ChaincodeContainerInfo
			result2 error
		})
	}
	fake.chaincodeContainerInfoReturnsOnCall[i] = struct {
		result1 *ccprovider.ChaincodeContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *Lifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.chaincodeDefinitionMutex.RLock()
	defer fake.chaincodeDefinitionMutex.RUnlock()
	fake.chaincodeContainerInfoMutex.RLock()
	defer fake.chaincodeContainerInfoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Lifecycle) recordInvocation(key string, args []interface{}) {
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