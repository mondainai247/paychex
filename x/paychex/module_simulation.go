package paychex

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"paychex/testutil/sample"
	paychexsimulation "paychex/x/paychex/simulation"
	"paychex/x/paychex/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = paychexsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateEmployee = "op_weight_msg_employee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateEmployee int = 100

	opWeightMsgUpdateEmployee = "op_weight_msg_employee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateEmployee int = 100

	opWeightMsgDeleteEmployee = "op_weight_msg_employee"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteEmployee int = 100

	opWeightMsgCreateQueryResponse = "op_weight_msg_query_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateQueryResponse int = 100

	opWeightMsgUpdateQueryResponse = "op_weight_msg_query_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateQueryResponse int = 100

	opWeightMsgDeleteQueryResponse = "op_weight_msg_query_response"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteQueryResponse int = 100

	opWeightMsgSendPayroll = "op_weight_msg_send_payroll"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendPayroll int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	paychexGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		EmployeeList: []types.Employee{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		EmployeeCount: 2,
		QueryResponseList: []types.QueryResponse{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		QueryResponseCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&paychexGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateEmployee int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateEmployee, &weightMsgCreateEmployee, nil,
		func(_ *rand.Rand) {
			weightMsgCreateEmployee = defaultWeightMsgCreateEmployee
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateEmployee,
		paychexsimulation.SimulateMsgCreateEmployee(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateEmployee int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateEmployee, &weightMsgUpdateEmployee, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateEmployee = defaultWeightMsgUpdateEmployee
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateEmployee,
		paychexsimulation.SimulateMsgUpdateEmployee(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteEmployee int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteEmployee, &weightMsgDeleteEmployee, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteEmployee = defaultWeightMsgDeleteEmployee
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteEmployee,
		paychexsimulation.SimulateMsgDeleteEmployee(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateQueryResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateQueryResponse, &weightMsgCreateQueryResponse, nil,
		func(_ *rand.Rand) {
			weightMsgCreateQueryResponse = defaultWeightMsgCreateQueryResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateQueryResponse,
		paychexsimulation.SimulateMsgCreateQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateQueryResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateQueryResponse, &weightMsgUpdateQueryResponse, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateQueryResponse = defaultWeightMsgUpdateQueryResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateQueryResponse,
		paychexsimulation.SimulateMsgUpdateQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteQueryResponse int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteQueryResponse, &weightMsgDeleteQueryResponse, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteQueryResponse = defaultWeightMsgDeleteQueryResponse
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteQueryResponse,
		paychexsimulation.SimulateMsgDeleteQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendPayroll int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendPayroll, &weightMsgSendPayroll, nil,
		func(_ *rand.Rand) {
			weightMsgSendPayroll = defaultWeightMsgSendPayroll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendPayroll,
		paychexsimulation.SimulateMsgSendPayroll(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateEmployee,
			defaultWeightMsgCreateEmployee,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgCreateEmployee(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateEmployee,
			defaultWeightMsgUpdateEmployee,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgUpdateEmployee(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteEmployee,
			defaultWeightMsgDeleteEmployee,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgDeleteEmployee(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateQueryResponse,
			defaultWeightMsgCreateQueryResponse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgCreateQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateQueryResponse,
			defaultWeightMsgUpdateQueryResponse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgUpdateQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteQueryResponse,
			defaultWeightMsgDeleteQueryResponse,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgDeleteQueryResponse(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendPayroll,
			defaultWeightMsgSendPayroll,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				paychexsimulation.SimulateMsgSendPayroll(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
