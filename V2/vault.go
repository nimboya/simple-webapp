package main

import (
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

var (
	vaultAddress = os.Getenv("VAULT_ADDR")
	token        = os.Getenv("VAULT_TOKEN")
	pathName     = "database/creds/vault-mysql-role"
)

func getVaultCreds() (interface{}, interface{}, error) {
	client, err := api.NewClient(&api.Config{
		Address: vaultAddress,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}

	client.SetToken(token)
	secretVaules, err := client.Logical().Read(pathName)
	if err != nil {
		log.Fatalf("%s", err)
	}

	sqlUser := secretVaules.Data["username"]
	sqlPass := secretVaules.Data["password"]

	return sqlUser, sqlPass, nil
}
