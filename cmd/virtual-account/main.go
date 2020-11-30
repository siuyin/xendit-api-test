package main

import (
	"fmt"
	"log"

	"github.com/siuyin/dflt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

func main() {
	fmt.Println("fixed virtual account creation example")

	xendit.Opt.SecretKey = dflt.EnvString("SecretKey", "mySecretKey")

	data := virtualaccount.CreateFixedVAParams{
		ExternalID: "demo-1475804036622",
		BankCode:   "BNI",
		Name:       "Rika Sutanto",
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created fixed va: %+v\n", resp)
}
