package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EmployeeList:      []Employee{},
		QueryResponseList: []QueryResponse{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in employee
	employeeIdMap := make(map[uint64]bool)
	employeeCount := gs.GetEmployeeCount()
	for _, elem := range gs.EmployeeList {
		if _, ok := employeeIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for employee")
		}
		if elem.Id >= employeeCount {
			return fmt.Errorf("employee id should be lower or equal than the last id")
		}
		employeeIdMap[elem.Id] = true
	}
	// Check for duplicated ID in queryResponse
	queryResponseIdMap := make(map[uint64]bool)
	queryResponseCount := gs.GetQueryResponseCount()
	for _, elem := range gs.QueryResponseList {
		if _, ok := queryResponseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for queryResponse")
		}
		if elem.Id >= queryResponseCount {
			return fmt.Errorf("queryResponse id should be lower or equal than the last id")
		}
		queryResponseIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
