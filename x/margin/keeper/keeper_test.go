package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sifchain/sifnode/x/margin/test"
	"github.com/Sifchain/sifnode/x/margin/types"
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

