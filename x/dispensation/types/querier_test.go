package types_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

const (
	QueryAllDistributions   = "distributions"
	QueryRecordsByDistrName = "records_by_name"
	QueryRecordsByRecipient = "records_by_recipient"
	QueryClaimsByType       = "claims_by_type"
)

func TestQueryRecordsByDistributionName_ValidBasic(t *testing.T) {
	distributionName := types.AttributeKeyDistributionName
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	t.Log(distributionName, status)
	result := types.NewQueryRecordsByDistributionName(distributionName, status)
	t.Log(result)
	assert.Equal(t, distributionName, result.DistributionName)

}

func TestQueryRecordsByRecipientAddr_ValidBasic(t *testing.T) {
	address := sdk.AccAddress("address___")
	t.Log(address.String())

	result := types.NewQueryRecordsByRecipientAddr(address.String())
	t.Log(result)
	assert.Equal(t, address.String(), result.Address)
}

func TestQueryUserClaims_ValidBasic(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	t.Log(distributiontype)
	result := types.NewQueryUserClaims(distributiontype)
	t.Log(result)
	assert.Equal(t, distributiontype, result.UserClaimType)
}
