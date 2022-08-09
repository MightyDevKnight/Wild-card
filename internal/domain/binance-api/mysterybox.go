package binance_api

type BuyRequest struct {
	Amount    int    `json:"number"`
	ProductID string `json:"productId"`
}

type NftMysteryBoxesInfoResponse struct {
	Data NftMysteryBoxInfoResponse `json:"data"`
}

type NftMysteryBoxInfoResponse struct {
	StartTime    int64  `json:"startTime"`
	Price        string `json:"price"`
	LimitPerTime uint64 `json:"limitPerTime"`
	UserBalance  string `json:"userBalance"`
}

