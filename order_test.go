package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
	"github.com/kamoney/sdk_golang/utility"
)

func TestListOrder(t *testing.T) {
	response, err := private.ListOrder(kamoney_sdk_dtos.ListOrderRequestParams{
		Begin:  "",
		End:    "",
		Search: "OKM92006666",
		Status: "CANCELED",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetOrderInfo(t *testing.T) {
	response, err := private.GetOrderInfo(kamoney_sdk_dtos.GetOrderInfoRequestParams{
		ID: "OKM92006666",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestListOrderReceipt(t *testing.T) {
	response, err := private.ListOrderReceipt(kamoney_sdk_dtos.ListOrderReceiptRequestParams{})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestGetOrderReceiptDownload(t *testing.T) {
	response, err := private.GetOrderReceiptDownload(kamoney_sdk_dtos.GetOrderReceiptDownloadRequestParams{
		Filename: "",
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}

func TestCreateOrder(t *testing.T) {
	response, err := private.CreateOrder(kamoney_sdk_dtos.CreateOrderRequestParams{
		Nonce:   fmt.Sprint(utility.GenNonce()),
		Asset:   "LTC",
		Network: "LTC",
		// PaymentSlips: []struct {
		// 	Barcode     string  `json:"barcode"`
		// 	Institution string  `json:"institution"`
		// 	Amount      float64 `json:"amount"`
		// 	DueDate     string  `json:"due_date"`
		// }{
		// 	{
		// 		Barcode:     "23793380296101222882356006333308191510000002000",
		// 		Institution: "Banco NuBank",
		// 		Amount:      100.50,
		// 		DueDate:     "2023-12-31",
		// 	},
		// },
		Pix: []struct {
			Type   string  `json:"type"`
			Key    string  `json:"key"`
			Amount float64 `json:"amount"`
		}{
			{
				Type:   "EMAIL",
				Key:    "igorasft@gmail.com",
				Amount: 100.50,
			},
		},
		// DigitalProducts: []struct {
		// 	ProductID int64 `json:"product_id"`
		// 	Quantity  int64 `json:"quantity"`
		// }{
		// 	{
		// 		ProductID: 138,
		// 		Quantity:  2,
		// 	},
		// },
		// DirectTransfers: []struct {
		// 	AccountType   string  `json:"account_type"`
		// 	BankID        int64   `json:"bank_id"`
		// 	Agency        string  `json:"agency"`
		// 	AccountNumber string  `json:"account_number"`
		// 	PersonalID    string  `json:"personal_id"`
		// 	Owner         string  `json:"owner"`
		// 	Amount        float64 `json:"amount"`
		// }{
		// 	{
		// 		AccountType:   "CC",
		// 		BankID:        341,
		// 		Agency:        "0001",
		// 		AccountNumber: "123456",
		// 		PersonalID:    "08830121622",
		// 		Owner:         "John Doe",
		// 		Amount:        100.50,
		// 	},
		// },
	})

	if err != nil {
		t.Error(err)
	}

	t.Log(response)
}
