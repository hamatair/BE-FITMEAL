package midtranss

import (
	"errors"
	"fmt"
	"intern-bcc/entity"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransServiceI interface {
	GenerateSnapUrl(t *entity.TopUp) error
	VerifyPayment(data map[string]interface{}) error
}

type Midtrans struct {
	Key    string
	IsProd bool
}

type MidtransServices struct {
	client         snap.Client
	MidtransConfig Midtrans
}

func NewMidtrans(cnf *Midtrans) MidtransServiceI {
	var client snap.Client
	envi := midtrans.Sandbox
	if cnf.IsProd {
		envi = midtrans.Production
	}

	cnf.Key = os.Getenv("MIDTRANS_KEY")

	client.New(cnf.Key, envi)

	return &MidtransServices{
		client:         client,
		MidtransConfig: *cnf,
	}
}

// GenerateSnapUrl implements MidtransService.
func (m *MidtransServices) GenerateSnapUrl(t *entity.TopUp) error {
	// currentTime := time.Now().UTC().Truncate(time.Nanosecond)
	// formattedTime := currentTime.Format("2006-01-02 15:04:05 -0700")
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  t.ID.String(),
			GrossAmt: int64(t.Amount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		Items: &[]midtrans.ItemDetails{
			{
				Name:  "FitMeal Premium",
				Price: 20000,
				Qty:   1,
			},
		},
		// Expiry: &snap.ExpiryDetails{
		// 	StartTime: formattedTime,
		// 	Unit:      "HOUR",
		// 	Duration:  int64(1 * time.Hour),
		// },
	}

	snaResp, err := m.client.CreateTransaction(req)
	t.SnapUrl = snaResp.RedirectURL

	if err != nil {
		return err
	}

	return nil
}

func (m *MidtransServices) VerifyPayment( data map[string]interface{}) error {
	var client coreapi.Client
	envi := midtrans.Sandbox
	if m.MidtransConfig.IsProd {
		envi = midtrans.Production
	}

	client.New(m.MidtransConfig.Key, envi)
	
	orderId, exists := data["order_id"].(string)
	if !exists {
		
		return errors.New("invalid payload")
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, err := client.CheckTransaction(orderId)
	if err != nil {
		return err
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					return nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
				fmt.Println("cancel")
			}

			return nil
		}
	}
	return nil
}
