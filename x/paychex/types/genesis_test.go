package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"paychex/x/paychex/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				EmployeeList: []types.Employee{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				EmployeeCount: 2,
				QueryResponseList: []types.QueryResponse{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				QueryResponseCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated employee",
			genState: &types.GenesisState{
				EmployeeList: []types.Employee{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid employee count",
			genState: &types.GenesisState{
				EmployeeList: []types.Employee{
					{
						Id: 1,
					},
				},
				EmployeeCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated queryResponse",
			genState: &types.GenesisState{
				QueryResponseList: []types.QueryResponse{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid queryResponse count",
			genState: &types.GenesisState{
				QueryResponseList: []types.QueryResponse{
					{
						Id: 1,
					},
				},
				QueryResponseCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
