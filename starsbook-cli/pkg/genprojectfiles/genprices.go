package genprojectfiles

func genPrices(nft NFTOutputJson, project ProjectOutputJson, lowestRarityScore float32, mintPrice int) PriceOutputJson {
	KRV := (nft.RarityScore / lowestRarityScore) * float32(mintPrice)
	KARV_NOW := KRV * 4.679
	KARV_FUTURE := KARV_NOW * 6.11

	KHRV := ((nft.RarityScore + float32(project.NumberOfNfts)) / (lowestRarityScore + float32(project.NumberOfNfts))) * float32(mintPrice)
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
