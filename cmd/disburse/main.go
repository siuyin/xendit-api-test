package main

import (
	"fmt"
	"log"
	"time"

	"github.com/siuyin/dflt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/disbursement"
)

func main() {
	fmt.Println("xendit api: disbursement test")

	xendit.Opt.SecretKey = dflt.EnvString("SecretKey", "mySecretKey")

	createData := disbursement.CreateParams{
		IdempotencyKey:    "disbursement" + time.Now().String(),
		ExternalID:        dflt.EnvString("ID", "456789"),
		BankCode:          "BCA",
		AccountHolderName: "RAIDY WIJAYA",
		AccountNumber:     "1234567890",
		Description:       "Disbursement from Go",
		Amount:            1000,
	}

	resp, err := disbursement.Create(&createData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created disbursement: %+v\n", resp)

}
