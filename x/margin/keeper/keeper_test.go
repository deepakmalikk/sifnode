package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	clptest "github.com/Sifchain/sifnode/x/clp/test"
	"github.com/Sifchain/sifnode/x/margin/test"
	"github.com/Sifchain/sifnode/x/margin/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_Errors(t *testing.T) {
	_, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
}

func TestKeeper_SetMTP(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	var mtp types.MTP
	marginKeeper.SetMTP(ctx, &mtp)
}

func TestKeeper_GetMTP(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.GetMTP(ctx, "xxx", "xxx")
}

func TestKeeper_GetMTPIterator(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.GetMTPIterator(ctx)
}

func TestKeeper_GetMTPs(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.GetMTPs(ctx)
}

func TestKeeper_GetMTPsForAsset(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.GetMTPsForAsset(ctx, "xxx")
}

func TestKeeper_GetAssetsForMTP(t *testing.T) {
	_, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
}

func TestKeeper_DestroyMTP(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.DestroyMTP(ctx, "xxx", "xxx")
}

func TestKeeper_ClpKeeper(t *testing.T) {
	_, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.ClpKeeper()
}

func TestKeeper_BankKeeper(t *testing.T) {
	_, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.BankKeeper()
}

func TestKeeper_GetLeverageParam(t *testing.T) {
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	marginKeeper.GetLeverageParam(ctx)
}

func TestKeeper_CustodySwap(t *testing.T) {
	pool := clptest.GenerateRandomPool(1)[0]
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	swapResult, err := marginKeeper.CustodySwap(ctx, pool, "xxx", sdk.NewUint(10000))
	assert.NotNil(t, swapResult)
	assert.Nil(t, err)
}

func TestKeeper_Borrow(t *testing.T) {
	pool := clptest.GenerateRandomPool(1)[0]
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	mtp := types.NewMTP()
	err := marginKeeper.Borrow(ctx, "xxx", sdk.NewUint(10000), sdk.NewUint(1000), mtp, pool, sdk.NewUint(1))
	assert.Nil(t, err)
}

func TestKeeper_UpdatePoolHealth(t *testing.T) {
	pool := clptest.GenerateRandomPool(1)[0]
	ctx, app := test.CreateTestAppMargin(false)
	marginKeeper := app.MarginKeeper
	assert.NotNil(t, marginKeeper)
	err := marginKeeper.UpdatePoolHealth(ctx, pool)
	assert.Nil(t, err)
}
