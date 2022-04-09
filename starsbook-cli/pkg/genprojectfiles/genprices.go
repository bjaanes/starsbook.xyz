package genprojectfiles

import (
	"fmt"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/attributes"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
)

func genPrices(nft NFTOutputJson, project conf.Project, attributeMap attributes.AttributeMap, lowestRarityScore float32, mintPrice int) PriceOutputJson {
	KRV := (nft.RarityScore / lowestRarityScore) * float32(mintPrice)
	KARV_NOW := KRV * 4.679
	KARV_FUTURE := KARV_NOW * 6.11

	if project.RarityOverrideAttribute != "" {
		var rarity string
		for _, attr := range nft.Attributes {
			if attr.Type == project.RarityOverrideAttribute {
				rarity = attr.Value
				break
			}
		}

		if rarity != "" {
			rarityCount := attributeMap[project.RarityOverrideAttribute][rarity]

			mostCommon := 0
			for _, n := range attributeMap[project.RarityOverrideAttribute] {
				if n > mostCommon {
					mostCommon = n
				}
			}

			KRV = float32(mintPrice) * (float32(mostCommon) / float32(project.NumberOfNFTs)) / (float32(rarityCount) / float32(project.NumberOfNFTs))
			fmt.Println(KRV, nft.Name, rarityCount)
			KARV_NOW = KRV * 4.679
			KARV_FUTURE = KARV_NOW * 6.11
		}
	}

	KHRV := ((nft.RarityScore + float32(project.NumberOfNFTs)) / (lowestRarityScore + float32(project.NumberOfNFTs))) * float32(mintPrice)
	KHARV_NOW := KHRV * 4.679
	KHARV_FUTURE := KHARV_NOW * 6.11

	return PriceOutputJson{
		KRV:         KRV,
		KARV_NOW:    KARV_NOW,
		KARV_FUTURE: KARV_FUTURE,
		KHRV:        KHRV,
		KHRV_NOW:    KHARV_NOW,
		KHRV_FUTURE: KHARV_FUTURE,
	}
}
