package response

import "encoding/json"

type OrderPayResponse struct {
	Head header `json:"head"`
	Body struct {
		TotalAmount    string `json:"totalAmount"`
		ClearDate      string `json:"clearDate"`
		Credential     string `json:"credential"`
		TradeNo        string `json:"tradeNo"`
		PayTime        string `json:"payTime"`
		BuyerPayAmount string `json:"buyerPayAmount"`
		OrderCode      string `json:"orderCode"`
		QrCode		   string	`json:"qrCode"`
		DiscAmount     string `json:"discAmount"`
		MerchExtendParams string `json:"merchExtendParams"`
	} `json:"body"`
}

func (resp *OrderPayResponse) SetData(data string) {
	dataByte := []byte(data)
	json.Unmarshal(dataByte, resp)
}
