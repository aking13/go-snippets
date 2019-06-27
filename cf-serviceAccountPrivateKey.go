package main

import (
	"context"
	"encoding/json"
	"log"

	"golang.org/x/oauth2/google"
)

type CredentialsFile struct {
	ClientEmail  string `json:"client_email"`
	ClientID     string `json:"client_id"`
	PrivateKey   string `json:"private_key"`
	PrivateKeyID string `json:"private_key_id"`
	ProjectID    string `json:"project_id"`
}

func DefaultCredentialsFile(ctx context.Context, scopes ...string) (*CredentialsFile, error) {
	creds, err := google.FindDefaultCredentials(ctx, scopes...)
	if err != nil {
		return nil, err
	}
	cf := new(CredentialsFile)
	if err := json.Unmarshal(creds.JSON, cf); err != nil {
		return nil, err
	}
	return cf, nil
}

func main() {
	ctx := context.Background()
	creds, err := DefaultCredentialsFile(ctx)
	if err != nil {
		log.Fatalf("Failed to find default credentials: %v", err)
	}
	log.Printf("creds: %#v\n", creds)
}
