package kamoney_sdk_test

import (
	"fmt"
	"testing"

	"github.com/kamoney/sdk_golang/kamoney_sdk_dtos"
)

func TestListOrder(t *testing.T) {
	response, err := private.ListOrder(kamoney_sdk_dtos.ListOrderRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetOrderInfo(t *testing.T) {
	response, err := private.GetOrderInfo(kamoney_sdk_dtos.GetOrderInfoRequestParams{
		ID: "OKM39113080",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestListOrderReceipt(t *testing.T) {
	response, err := private.ListOrderReceipt(kamoney_sdk_dtos.ListOrderReceiptRequestParams{})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestGetOrderReceiptDownload(t *testing.T) {
	response, err := private.GetOrderReceiptDownload(kamoney_sdk_dtos.GetOrderReceiptDownloadRequestParams{
		Filename: "",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

func TestCreateOrder(t *testing.T) {
	response, err := private.CreateOrder(kamoney_sdk_dtos.CreateOrderRequestParams{
		Asset:   "LTC",
		Network: "LTC",
		PaymentSlips: []struct {
			Barcode     string  `json:"barcode"`
			Institution string  `json:"institution"`
			Amount      float64 `json:"amount"`
			DueDate     string  `json:"due_date"`
		}{
			{
				Barcode:     "23793380296101222882356006333308191510000002000",
				Institution: "Banco NuBank",
				Amount:      100.50,
				DueDate:     "2023-12-31",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}
