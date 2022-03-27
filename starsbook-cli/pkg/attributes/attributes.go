package attributes

import (
	"github.com/go-errors/errors"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/conf"
	"github.com/starsbook/starsbook.xyz/starsbook-cli/pkg/nftinfo"
)

type NumberOfInstancesOfTrait map[string]int          // Attribute value -> number of nfts with that value
type AttributeMap map[string]NumberOfInstancesOfTrait // Attribute type -> map of values

func GenerateRarityAttributeMap(p conf.Project) (AttributeMap, error) {
	nfts, err := nftinfo.GetNFTInfos(p)
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	attributeMap := make(AttributeMap)
	for _, nft := range nfts {
		for _, attr := range nft.Attributes {
			if IgnoreForRarity(p, attr) {
				continue
			}

			if attributeMap[attr.TraitType] == nil {
				attributeMap[attr.TraitType] = make(NumberOfInstancesOfTrait)
			}

			value := attr.GetValue()
			attributeMap[attr.TraitType][value] = attributeMap[attr.TraitType][value] + 1
		}
	}

	return attributeMap, nil
}

func IgnoreForRarity(p conf.Project, attr nftinfo.Attribute) bool {
	for _, ign := range p.AttributesToIgnoreForRarity {
		if ign == attr.TraitType {
			return true
		}
	}

	return false
}

func IgnoreForDisplay(p conf.Project, attr nftinfo.Attribute) bool {
	for _, ign := range p.AttributesToIgnoreForDisplay {
		if ign == attr.TraitType {
			return true
		}
	}

	return false
}
