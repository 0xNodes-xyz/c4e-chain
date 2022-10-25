package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	testapp "github.com/chain4energy/c4e-chain/testutil/app"
	commontestutils "github.com/chain4energy/c4e-chain/testutil/common"
)

func TestSendVestingAccount(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt(), accInitBalance.Sub(vested), vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, 1, sdk.NewInt(100), true)

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

// TODO add test with restart vesting set to false

func TestSendVestingAccountVestingPoolNotExistsForAddress(t *testing.T) {
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, 2, sdk.NewInt(100), true,
		"send to new vesting account - withdraw all available error: withdraw all available - no vesting pools found error: address: cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq: not found")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountVestingPoolNotFound(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, 2, sdk.NewInt(100), true, "send to new vesting account - vesting pool with id 2 not found: not found")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccounNotEnoughToSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, 1, sdk.NewInt(1100), true, "send to new vesting account - vesting available: 1000 is smaller than requested amount: 1100: insufficient funds")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountNotEnoughToSendAferSuccesfulSend(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)
	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, 1, sdk.NewInt(100), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, 1, sdk.NewInt(950), true, "send to new vesting account - vesting available: 900 is smaller than requested amount: 950: insufficient funds")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}

func TestSendVestingAccountAlreadyExists(t *testing.T) {
	vested := sdk.NewInt(1000)
	testHelper := testapp.SetupTestAppWithHeight(t, 1000)

	acountsAddresses, _ := commontestutils.CreateAccounts(2, 0)

	accAddr := acountsAddresses[0]
	accAddr2 := acountsAddresses[1]

	accInitBalance := sdk.NewInt(10000)
	testHelper.BankUtils.AddDefaultDenomCoinsToAccount(accInitBalance, accAddr)

	vestingTypes := testHelper.C4eVestingUtils.SetupVestingTypes(2, 1, 1)
	usedVestingType := vestingTypes.VestingTypes[0]

	testHelper.C4eVestingUtils.MessageCreateVestingPool(accAddr, false, true, vPool1, 1000, *usedVestingType, vested, accInitBalance, sdk.ZeroInt() /*0,*/, accInitBalance.Sub(vested) /*0,*/, vested)

	testHelper.C4eVestingUtils.MessageSendToVestingAccount(accAddr, accAddr2, 1, sdk.NewInt(100), true)
	testHelper.C4eVestingUtils.MessageSendToVestingAccountError(accAddr, accAddr2, 1, sdk.NewInt(300), true, "new vesting account - account address: "+accAddr2.String()+": entity already exists")

	testHelper.C4eVestingUtils.ValidateGenesisAndInvariants()
}
