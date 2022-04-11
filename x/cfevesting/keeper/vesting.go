package keeper

import (
	"math"
	"strconv"
	"time"

	"github.com/chain4energy/c4e-chain/x/cfevesting/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// vest

func (k Keeper) Vest(ctx sdk.Context, addr string, amount sdk.Int, vestingType string) error {
	k.Logger(ctx).Debug("Vest: addr: " + addr + "amount: " + amount.String() + "vestingType: " + vestingType)
	vt, err := k.GetVestingType(ctx, vestingType)
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())

		return sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
	}
	k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))

	// return k.addVesting(ctx, false, addr, addr, amount, vestingType, vt.DelegationsAllowed, ctx.BlockHeight(),
	// 	vt.LockupPeriod.Nanoseconds()+ctx.BlockHeight(), vt.LockupPeriod.Nanoseconds()+vt.VestingPeriod.Nanoseconds()+ctx.BlockHeight(),
	// 	vt.TokenReleasingPeriod.Nanoseconds())

	return k.addVesting(ctx, false, addr, addr, amount, vestingType, vt.DelegationsAllowed, ctx.BlockTime(),
		ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod),
		vt.TokenReleasingPeriod)

}

func (k Keeper) addVesting(
	ctx sdk.Context,
	isModuleInternalOperation bool,
	vestingAddr string,
	coinSrcAddr string,
	amount sdk.Int,
	vestingType string,
	delegationAllowed bool,
	vestingStart time.Time,
	lockEnd time.Time,
	vestingEnd time.Time,
	releasePeriod time.Duration) error {

	if amount.Equal(sdk.ZeroInt()) {
		return nil
	}

	vestingAccAddress, err := sdk.AccAddressFromBech32(vestingAddr)
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}

	denom := k.GetParams(ctx).Denom
	k.Logger(ctx).Debug("denom: " + denom)

	var srcAccAddress sdk.AccAddress
	if !isModuleInternalOperation || delegationAllowed {
		srcAccAddress, err = sdk.AccAddressFromBech32(coinSrcAddr)
		if err != nil {
			k.Logger(ctx).Error("Error: " + err.Error())
			return err
		}
		k.Logger(ctx).Debug("vestingAccAddress: " + vestingAccAddress.String())

		balance := k.bank.GetBalance(ctx, srcAccAddress, denom)
		k.Logger(ctx).Debug("balance: " + balance.Amount.String())

		if balance.Amount.LT(amount) {
			k.Logger(ctx).Error("Error: " + "Balance [" + balance.Amount.String() +
				balance.Denom + "]lesser than requested amount: " + amount.String() + denom)
			return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Balance ["+balance.Amount.String()+
				balance.Denom+"]lesser than requested amount: "+amount.String()+denom)
		}
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, vestingAddr)
	k.Logger(ctx).Debug("vestingsFound: " + strconv.FormatBool(vestingsFound))
	var id int32
	var delegatableAddress sdk.AccAddress
	if !vestingsFound {
		accVestings = types.AccountVestings{}
		accVestings.Address = vestingAddr
		k.Logger(ctx).Debug("accVestings.Address: " + accVestings.Address)

		if delegationAllowed {
			delegatableAddress = k.CreateModuleAccountForDelegatableVesting(ctx, vestingAddr).GetAddress()
			k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())

			accVestings.DelegableAddress = delegatableAddress.String()
		}
		id = 1
	} else {
		if delegationAllowed {
			if accVestings.DelegableAddress == "" {
				delegatableAddress = k.CreateModuleAccountForDelegatableVesting(ctx, vestingAddr).GetAddress()
				k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())
				accVestings.DelegableAddress = delegatableAddress.String()
			} else {
				delegatableAddress, err = sdk.AccAddressFromBech32(accVestings.DelegableAddress)
				k.Logger(ctx).Debug("delegatableAddress: " + delegatableAddress.String())
				if err != nil {
					k.Logger(ctx).Error("Error: " + err.Error())
					return err
				}
			}
		}
		id = int32(len(accVestings.Vestings)) + 1
	}

	vesting := types.Vesting{
		Id:                        id,
		VestingType:               vestingType,
		VestingStart:              vestingStart,
		LockEnd:                   lockEnd,
		VestingEnd:                vestingEnd,
		Vested:                    amount,
		ReleasePeriod:             releasePeriod,
		DelegationAllowed:         delegationAllowed,
		Withdrawn:                 sdk.ZeroInt(),
		Sent:                      sdk.ZeroInt(),
		LastModification:          vestingStart,
		LastModificationVested:    amount,
		LastModificationWithdrawn: sdk.ZeroInt(),
	}
	accVestings.Vestings = append(accVestings.Vestings, &vesting)

	coinToSend := sdk.NewCoin(denom, amount)
	coinsToSend := sdk.NewCoins(coinToSend)

	if delegationAllowed {

		err = k.bank.SendCoins(ctx, srcAccAddress, delegatableAddress, coinsToSend)
		k.Logger(ctx).Info("after SendCoins: " + coinToSend.Amount.String())

	} else if !isModuleInternalOperation {
		err = k.bank.SendCoinsFromAccountToModule(ctx, srcAccAddress, types.ModuleName, coinsToSend)
		k.Logger(ctx).Info("after SendCoinsFromAccountToModule: " + coinToSend.Amount.String())

	}
	if err != nil {
		k.Logger(ctx).Error("Error: " + err.Error())
		return err
	}
	k.SetAccountVestings(ctx, accVestings)
	k.Logger(ctx).Info("Vest exit: " + vestingAddr)
	return nil
}

func (k Keeper) SendVesting(ctx sdk.Context, fromAddr string, toAddr string, vestingId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	if fromAddr == toAddr {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.Vestings) == 0 {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vestings found")
	}
	var vesting *types.Vesting = nil
	for _, vest := range accVestings.Vestings {
		if vest.Id == vestingId {
			vesting = vest
		}
	}
	if vesting == nil {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" not found")
	}
	if !vesting.TransferAllowed {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" is not tranferable")
	}
	available := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	if available.LT(amount) {
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"vesting available: %s is smaller than %s", available, amount)
	}
	if vesting.DelegationAllowed {
		if len(accVestings.DelegableAddress) == 0 {
			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrLogic, "delegable vesting has no delegable address")
		}
		denom := k.GetParams(ctx).Denom
		from, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		if err != nil {
			return withdrawn, err
		}
		balance := k.bank.GetBalance(ctx, from, denom)
		lockedCoins := k.bank.LockedCoins(ctx, from)
		locked := sdk.NewCoin(denom, lockedCoins.AmountOf(denom))
		spendable := balance.Sub(locked)
		if spendable.Amount.LT(amount) {
			return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
				"vesting available: %s is smaller than %s - probably delageted to validator.", spendable.Amount, amount)
		}
	}
	vesting.Sent = amount
	vesting.LastModification = ctx.BlockTime()
	vesting.LastModificationVested = available.Sub(amount)
	vesting.LastModificationWithdrawn = sdk.ZeroInt()

	if restartVesting {
		vt, vErr := k.GetVestingType(ctx, vesting.VestingType)
		if vErr != nil {
			k.Logger(ctx).Error("Error: " + err.Error())

			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))
		// return w, k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, ctx.BlockHeight(),
		// 	vt.LockupPeriod.Nanoseconds()+ctx.BlockHeight(), vt.LockupPeriod.Nanoseconds()+vt.VestingPeriod.Nanoseconds()+ctx.BlockHeight(),
		// 	vt.TokenReleasingPeriod.Nanoseconds())
		err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, /*TODO iit should be checks vesting aprameter, but also dst shoul by like in type*/
			ctx.BlockTime(),
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod),
			vt.TokenReleasingPeriod)
	} else {
		err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vesting.DelegationAllowed, ctx.BlockTime(),
			vesting.LockEnd, vesting.VestingEnd,
			vesting.ReleasePeriod)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings)
	}
	return w, err
}

func (k Keeper) WithdrawAllAvailable(ctx sdk.Context, addr string) (withdrawn sdk.Coin, returnedError error) {
	k.Logger(ctx).Debug("WithdrawAllAvailable: " + addr)

	accAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, addr)
	if !vestingsFound {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	if len(accVestings.Vestings) == 0 {
		return withdrawn, status.Error(codes.NotFound, "No vestings")
	}

	current := ctx.BlockTime()
	toWithdraw := sdk.ZeroInt()
	toWithdrawDelegable := sdk.ZeroInt()
	toWithdrawDelegableMap := make(map[int32]sdk.Int)
	for _, vesting := range accVestings.Vestings {
		withdrawable := CalculateWithdrawable(current, *vesting)
		// vesting.Withdrawn = vesting.Withdrawn.Add(withdrawable)
		// vesting.LastModificationWithdrawn = vesting.LastModificationWithdrawn.Add(withdrawable)
		if vesting.DelegationAllowed {
			toWithdrawDelegableMap[vesting.Id] = withdrawable
			toWithdrawDelegable = toWithdrawDelegable.Add(withdrawable)
		} else {
			vesting.Withdrawn = vesting.Withdrawn.Add(withdrawable)
			vesting.LastModificationWithdrawn = vesting.LastModificationWithdrawn.Add(withdrawable)
			toWithdraw = toWithdraw.Add(withdrawable)
		}
	}

	denom := k.GetParams(ctx).Denom
	if toWithdraw.GT(sdk.ZeroInt()) {
		coinToSend := sdk.NewCoin(denom, toWithdraw)
		coinsToSend := sdk.NewCoins(coinToSend)
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, accAddress, coinsToSend)
		if err != nil {
			return withdrawn, err
		}
	}
	withdrawnDelegable := toWithdrawDelegable
	if toWithdrawDelegable.GT(sdk.ZeroInt()) {
		delegatableAddress, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		if err != nil {
			return withdrawn, err
		}
		coinToSend := sdk.NewCoin(denom, toWithdrawDelegable)
		available := k.bank.GetBalance(ctx, delegatableAddress, denom)
		if available.Amount.GT(sdk.ZeroInt()) {
			if available.Amount.LT(coinToSend.Amount) {
				coinToSend = available
				withdrawnDelegable = coinToSend.Amount
			}
			coinsToSend := sdk.NewCoins(coinToSend)

			err = k.bank.SendCoins(ctx, delegatableAddress, accAddress, coinsToSend)
			if err != nil {
				return withdrawn, err
			}
		} else {
			withdrawnDelegable = sdk.ZeroInt()
		}
	}

	withdrawnDelegableHelper := withdrawnDelegable
	for _, vesting := range accVestings.Vestings {
		if vesting.DelegationAllowed {
			toWithdraw := toWithdrawDelegableMap[vesting.Id]
			diff := withdrawnDelegableHelper.Sub(toWithdraw)
			if diff.IsNegative() {
				toWithdraw = toWithdraw.Add(diff)
			}
			withdrawnDelegableHelper = withdrawnDelegableHelper.Sub(toWithdraw)
			vesting.Withdrawn = vesting.Withdrawn.Add(toWithdraw)
			vesting.LastModificationWithdrawn = vesting.LastModificationWithdrawn.Add(toWithdraw)
		}
	}

	k.SetAccountVestings(ctx, accVestings)

	return sdk.NewCoin(denom, toWithdraw.Add(withdrawnDelegable)), nil
}

func (k Keeper) SendToNewVestingAccount(ctx sdk.Context, fromAddr string, toAddr string, vestingId int32, amount sdk.Int, restartVesting bool) (withdrawn sdk.Coin, returnedError error) {
	if fromAddr == toAddr {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "from address and to address cannot be identical")
	}

	w, err := k.WithdrawAllAvailable(ctx, fromAddr)
	if err != nil {
		return withdrawn, err
	}

	accVestings, vestingsFound := k.GetAccountVestings(ctx, fromAddr)
	if !vestingsFound || len(accVestings.Vestings) == 0 {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "no vestings found")
	}
	var vesting *types.Vesting = nil
	for _, vest := range accVestings.Vestings {
		if vest.Id == vestingId {
			vesting = vest
		}
	}
	if vesting == nil {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" not found")
	}
	if !vesting.TransferAllowed {
		return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "vesting with id "+strconv.FormatInt(int64(vestingId), 10)+" is not tranferable")
	}
	available := vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	if available.LT(amount) {
		return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
			"vesting available: %s is smaller than %s", available, amount)
	}
	if vesting.DelegationAllowed {
		if len(accVestings.DelegableAddress) == 0 {
			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrLogic, "delegable vesting has no delegable address")
		}
		denom := k.GetParams(ctx).Denom
		from, err := sdk.AccAddressFromBech32(accVestings.DelegableAddress)
		if err != nil {
			return withdrawn, err
		}
		balance := k.bank.GetBalance(ctx, from, denom)
		lockedCoins := k.bank.LockedCoins(ctx, from)
		locked := sdk.NewCoin(denom, lockedCoins.AmountOf(denom))
		spendable := balance.Sub(locked)
		if spendable.Amount.LT(amount) {
			return withdrawn, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
				"vesting available: %s is smaller than %s - probably delageted to validator.", spendable.Amount, amount)
		}
	}
	vesting.Sent = amount
	vesting.LastModification = ctx.BlockTime()
	vesting.LastModificationVested = available.Sub(amount)
	vesting.LastModificationWithdrawn = sdk.ZeroInt()

	if restartVesting {
		vt, vErr := k.GetVestingType(ctx, vesting.VestingType)
		if vErr != nil {
			k.Logger(ctx).Error("Error: " + err.Error())

			return withdrawn, sdkerrors.Wrap(sdkerrors.ErrNotFound, err.Error())
		}
		k.Logger(ctx).Debug("vt: DelegationsAllowed: " + strconv.FormatBool(vt.DelegationsAllowed))
		// return w, k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, ctx.BlockHeight(),
		// 	vt.LockupPeriod.Nanoseconds()+ctx.BlockHeight(), vt.LockupPeriod.Nanoseconds()+vt.VestingPeriod.Nanoseconds()+ctx.BlockHeight(),
		// 	vt.TokenReleasingPeriod.Nanoseconds())
		// err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vt.DelegationsAllowed, ctx.BlockTime(),
		// 	ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod),
		// 	vt.TokenReleasingPeriod)

		err = k.createVestingAccount(ctx, vesting.DelegationAllowed, accVestings.DelegableAddress, toAddr, amount,
			ctx.BlockTime().Add(vt.LockupPeriod), ctx.BlockTime().Add(vt.LockupPeriod).Add(vt.VestingPeriod))
	} else {
		err = k.addVesting(ctx, true, toAddr, accVestings.DelegableAddress, amount, vesting.VestingType, vesting.DelegationAllowed, ctx.BlockTime(),
			vesting.LockEnd, vesting.VestingEnd,
			vesting.ReleasePeriod)

		err = k.createVestingAccount(ctx, vesting.DelegationAllowed, accVestings.DelegableAddress, toAddr, amount,
			vesting.LockEnd, vesting.VestingEnd)
	}
	if err == nil {
		k.SetAccountVestings(ctx, accVestings)
	}
	return w, err
}

func (k Keeper) CreateVestingAccount(ctx sdk.Context, fromAddress string, toAddress string,
	amount sdk.Coins, startTime int64, endTime int64) error {
	ak := k.account
	bk := k.bank

	if err := bk.IsSendEnabledCoins(ctx, amount...); err != nil {
		return err
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return err
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return err
	}

	if bk.BlockedAddr(to) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), amount.Sort(), endTime)

	var acc authtypes.AccountI

	acc = vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime)

	ak.SetAccount(ctx, acc)

	// defer func() {
	// 	telemetry.IncrCounter(1, "new", "account")

	// 	for _, a := range msg.Amount {
	// 		if a.Amount.IsInt64() {
	// 			telemetry.SetGaugeWithLabels(
	// 				[]string{"tx", "msg", "create_vesting_account"},
	// 				float32(a.Amount.Int64()),
	// 				[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
	// 			)
	// 		}
	// 	}
	// }()

	err = bk.SendCoins(ctx, from, to, amount)
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return nil
}

func CalculateWithdrawable(current time.Time, vesting types.Vesting) sdk.Int {
	if current.Equal(vesting.VestingStart) || current.Before(vesting.VestingStart) {
		return sdk.ZeroInt()
	}
	if current.Equal(vesting.LockEnd) || current.Before(vesting.LockEnd) {
		return sdk.ZeroInt()
	}
	if current.Equal(vesting.VestingEnd) || current.After(vesting.VestingEnd) {
		return vesting.LastModificationVested.Sub(vesting.LastModificationWithdrawn)
	}

	var lockEnd time.Time
	if vesting.VestingStart.After(vesting.LockEnd) {
		lockEnd = vesting.VestingStart
	} else {
		lockEnd = vesting.LockEnd
	}
	if vesting.GetLastModification().After(lockEnd) {
		lockEnd = vesting.LastModification
	}
	wholeVestingPariod := vesting.VestingEnd.Sub(lockEnd).Nanoseconds()
	fromStart := current.Sub(lockEnd).Nanoseconds()
	numOfPeriodsFromStart := fromStart / vesting.ReleasePeriod.Nanoseconds()
	numOfPeriods := int64(math.Ceil(float64(wholeVestingPariod) / float64(vesting.ReleasePeriod)))

	vested := vesting.LastModificationVested

	withdrawableFromStart := vested.MulRaw(numOfPeriodsFromStart).QuoRaw(numOfPeriods)
	return withdrawableFromStart.Sub(vesting.LastModificationWithdrawn)

}

func (k Keeper) CreateModuleAccountForDelegatableVesting(ctx sdk.Context, address string) authtypes.ModuleAccountI {
	perms := []string{}
	macc := authtypes.NewEmptyModuleAccount(createAccounNameForDelegatableVesting(address), perms...)
	maccI := (k.account.NewAccount(ctx, macc)).(authtypes.ModuleAccountI) // set the account number
	k.account.SetModuleAccount(ctx, maccI)
	return maccI
}

func (k Keeper) GetModuleAccountForDelegableVesting(ctx sdk.Context, address sdk.AccAddress) authtypes.ModuleAccountI {
	maccI := k.account.GetAccount(ctx, address).(authtypes.ModuleAccountI)
	return maccI
}

func createAccounNameForDelegatableVesting(address string) string {
	return types.ModuleName + address
}

func (k Keeper) createVestingAccount(ctx sdk.Context, delegationAllowed bool, fromAddress string, toAddress string, amount sdk.Int,
	lockEnd time.Time,
	vestingEnd time.Time) error {
	ak := k.account
	bk := k.bank

	denom := k.GetParams(ctx).Denom
	coinToSend := sdk.NewCoin(denom, amount)
	if err := bk.IsSendEnabledCoins(ctx, coinToSend); err != nil {
		return err
	}

	from, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return err
	}
	to, err := sdk.AccAddressFromBech32(toAddress)
	if err != nil {
		return err
	}

	if bk.BlockedAddr(to) {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", toAddress)
	}

	if acc := ak.GetAccount(ctx, to); acc != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "account %s already exists", toAddress)
	}

	baseAccount := ak.NewAccountWithAddress(ctx, to)
	if _, ok := baseAccount.(*authtypes.BaseAccount); !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid account type; expected: BaseAccount, got: %T", baseAccount)
	}
	coinsToSend := sdk.NewCoins(coinToSend)

	baseVestingAccount := vestingtypes.NewBaseVestingAccount(baseAccount.(*authtypes.BaseAccount), coinsToSend.Sort(), vestingEnd.Unix())

	var acc authtypes.AccountI

	startTime := lockEnd
	if lockEnd.Before(ctx.BlockTime()) {
		startTime = ctx.BlockTime()
	}
	// if msg.Delayed {
	// 	acc = types.NewDelayedVestingAccountRaw(baseVestingAccount)
	// } else {
	acc = vestingtypes.NewContinuousVestingAccountRaw(baseVestingAccount, startTime.Unix())
	// }

	ak.SetAccount(ctx, acc)

	// defer func() {
	// 	telemetry.IncrCounter(1, "new", "account")

	// 	for _, a := range msg.Amount {
	// 		if a.Amount.IsInt64() {
	// 			telemetry.SetGaugeWithLabels(
	// 				[]string{"tx", "msg", "create_vesting_account"},
	// 				float32(a.Amount.Int64()),
	// 				[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
	// 			)
	// 		}
	// 	}
	// }()

	// err = bk.SendCoins(ctx, from, to, coinsToSend)
	// if err != nil {
	// 	return err
	// }

	if delegationAllowed {

		err = bk.SendCoins(ctx, from, to, coinsToSend)
		k.Logger(ctx).Info("after SendCoins: " + coinToSend.Amount.String())

	} else {
		err = k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, coinsToSend)
		k.Logger(ctx).Info("after SendCoinsFromModuleToAccount: " + coinToSend.Amount.String())

	}
	if err != nil {
		return err
	}
	// ctx.EventManager().EmitEvent(
	// 	sdk.NewEvent(
	// 		sdk.EventTypeMessage,
	// 		sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
	// 	),
	// )

	return nil
}
