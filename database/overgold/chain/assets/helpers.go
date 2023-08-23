package assets

import (
	assetstypes "git.ooo.ua/vipcoin/chain/x/assets/types"
	extratypes "git.ooo.ua/vipcoin/chain/x/types"
	"github.com/lib/pq"

	"github.com/forbole/bdjuno/v3/database/types"
)

const (
	tableAssets         = "overgold_chain_assets_assets"
	tableCreateAssets   = "overgold_chain_assets_create"
	tableManageAsset    = "overgold_chain_assets_manage"
	tableSetExtrasAsset = "overgold_chain_assets_set_extra"
)

// toExtrasDB - mapping func to database model
func toExtrasDB(extras []*extratypes.Extra) types.ExtraDB {
	result := make([]extratypes.Extra, 0, len(extras))
	for _, extra := range extras {
		result = append(result, *extra)
	}

	return types.ExtraDB{Extras: result}
}

// fromExtrasDB - mapping func from database model
func fromExtrasDB(extras types.ExtraDB) []*extratypes.Extra {
	result := make([]*extratypes.Extra, 0, len(extras.Extras))
	for i := range extras.Extras {
		result = append(result, &extras.Extras[i])
	}

	return result
}

// toPoliciesDB - mapping func to database model
func toPoliciesDB(policies []assetstypes.AssetPolicy) pq.Int32Array {
	result := make(pq.Int32Array, 0, len(policies))
	for _, policy := range policies {
		result = append(result, assetstypes.AssetPolicy_value[policy.String()])
	}

	return result
}

// toPoliciesDomain - mapping func to domain model
func toPoliciesDomain(policies pq.Int32Array) []assetstypes.AssetPolicy {
	result := make([]assetstypes.AssetPolicy, 0, len(policies))
	for _, policy := range policies {
		result = append(result, assetstypes.AssetPolicy(policy))
	}

	return result
}

// toAssetDatabase - mapping func to database model
func toAssetDatabase(assets *assetstypes.Asset) types.DBAssets {
	return types.DBAssets{
		Issuer:        assets.Issuer,
		Name:          assets.Name,
		Policies:      toPoliciesDB(assets.Policies),
		State:         int32(assets.State),
		Issued:        assets.Issued,
		Burned:        assets.Burned,
		Withdrawn:     assets.Withdrawn,
		InCirculation: assets.InCirculation,
		Precision:     assets.Properties.Precision,
		FeePercent:    assets.Properties.FeePercent,
		Extras:        toExtrasDB(assets.Extras),
	}
}

// toAssetsArrDatabase- mapping func to database model
func toAssetsArrDatabase(assets ...*assetstypes.Asset) []types.DBAssets {
	result := make([]types.DBAssets, 0, len(assets))
	for _, asset := range assets {
		result = append(result, toAssetDatabase(asset))
	}

	return result
}

// toCreateAssetDatabase - mapping func to database model
func toCreateAssetDatabase(msg *assetstypes.MsgAssetCreate, transactionHash string) types.DBAssetCreate {
	return types.DBAssetCreate{
		Hash:       transactionHash,
		Creator:    msg.Creator,
		Name:       msg.Name,
		Issuer:     msg.Issuer,
		Policies:   toPoliciesDB(msg.Policies),
		State:      int32(msg.State),
		Precision:  msg.Properties.Precision,
		FeePercent: msg.Properties.FeePercent,
		Extras:     toExtrasDB(msg.Extras),
	}
}

// toCreateAssetDomain - mapping func from database model
func toCreateAssetDomain(asset types.DBAssetCreate) *assetstypes.MsgAssetCreate {
	return &assetstypes.MsgAssetCreate{
		Creator:  asset.Creator,
		Name:     asset.Name,
		Issuer:   asset.Issuer,
		Policies: toPoliciesDomain(asset.Policies),
		State:    assetstypes.AssetState(asset.State),
		Properties: assetstypes.Properties{
			Precision:  asset.Precision,
			FeePercent: asset.FeePercent,
		},
		Extras: fromExtrasDB(asset.Extras),
	}
}

// toManageAssetDatabase - mapping func to database model
func toManageAssetDatabase(msg *assetstypes.MsgAssetManage, transactionHash string) types.DBAssetManage {
	return types.DBAssetManage{
		Hash:          transactionHash,
		Creator:       msg.Creator,
		Name:          msg.Name,
		Policies:      toPoliciesDB(msg.Policies),
		State:         int32(msg.State),
		Issued:        msg.Issued,
		Burned:        msg.Burned,
		Withdrawn:     msg.Withdrawn,
		InCirculation: msg.InCirculation,
		Precision:     msg.Properties.Precision,
		FeePercent:    msg.Properties.FeePercent,
	}
}

// toManageAssetDomain - mapping func from database model
func toManageAssetDomain(asset types.DBAssetManage) *assetstypes.MsgAssetManage {
	return &assetstypes.MsgAssetManage{
		Creator:  asset.Creator,
		Name:     asset.Name,
		Policies: toPoliciesDomain(asset.Policies),
		State:    assetstypes.AssetState(asset.State),
		Properties: assetstypes.Properties{
			Precision:  asset.Precision,
			FeePercent: asset.FeePercent,
		},
		Issued:        asset.Issued,
		Burned:        asset.Burned,
		Withdrawn:     asset.Withdrawn,
		InCirculation: asset.InCirculation,
	}
}

// toAssetDomain - mapping func to domain model
func toAssetDomain(asset types.DBAssets) *assetstypes.Asset {
	return &assetstypes.Asset{
		Issuer:        asset.Issuer,
		Name:          asset.Name,
		Policies:      toPoliciesDomain(asset.Policies),
		State:         assetstypes.AssetState(asset.State),
		Issued:        asset.Issued,
		Burned:        asset.Burned,
		Withdrawn:     asset.Withdrawn,
		InCirculation: asset.InCirculation,
		Properties: assetstypes.Properties{
			Precision:  asset.Precision,
			FeePercent: asset.FeePercent,
		},
		Extras: fromExtrasDB(asset.Extras),
	}
}

// toSetExtraDatabase - mapping func to database model
func toSetExtraDatabase(msg *assetstypes.MsgAssetSetExtra, transactionHash string) types.DBAssetSetExtra {
	return types.DBAssetSetExtra{
		Hash:    transactionHash,
		Creator: msg.Creator,
		Name:    msg.Name,
		Extras:  toExtrasDB(msg.Extras),
	}
}

// toSetExtraDomain - mapping func to database model
func toSetExtraDomain(msg types.DBAssetSetExtra) *assetstypes.MsgAssetSetExtra {
	return &assetstypes.MsgAssetSetExtra{
		Creator: msg.Creator,
		Name:    msg.Name,
		Extras:  fromExtrasDB(msg.Extras),
	}
}
