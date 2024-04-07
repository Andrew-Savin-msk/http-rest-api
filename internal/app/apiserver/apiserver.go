package apiserver

import (
	"crypto/ecdsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"log"
	"net/http"
	"os"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.NewStore(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getPublicJWTKey(KeyURL string) *ecdsa.PublicKey {
	keyBytes, err := os.ReadFile(KeyURL)
	if err != nil {
		log.Fatalf("Unable to read PEM file: %v", err)
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatal("Failed to decode PEM block containing the key")
	}

	// Преобразование DER в публичный ключ
	pubKeyGeneric, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("Unable to parse ECDSA public key: %v", err)
	}

	pubKey, ok := pubKeyGeneric.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Not an ECDSA public key")
	}

	return pubKey
}
