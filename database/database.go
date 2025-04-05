package database

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	mysql2 "github.com/go-sql-driver/mysql"
	"os"
	"testify/models"
	"testify/seed"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBConn() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")
	sslMode := os.Getenv("DB_TLS_MODE")
	sslCertPath := os.Getenv("DB_TLS_CERT_PATH")

	// Load TLS certificate if required
	if sslMode == "verify-ca" || sslMode == "verify-full" {
		err := mysql2.RegisterTLSConfig("rds", &tls.Config{
			RootCAs: mustLoadCert(sslCertPath),
		})
		if err != nil {
			fmt.Println("Error registering TLS config: ", err)
			return nil, err
		}
		dsn += "&tls=rds"
	}

	mySqlConfig := mysql.Config{
		DriverName: os.Getenv("DRIVER_NAME"),
		DSN:        dsn,
	}

	db, err := gorm.Open(mysql.New(mySqlConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil, err
	}

	if err = db.AutoMigrate(
		&models.User{},
		&models.Question{},
		&models.Answer{},
		&models.UserTestSession{},
		&models.UserResponse{},
	); err != nil {
		fmt.Println("Error migrating models: ", err)
		return nil, err
	}

	seed.SeedAll(db)

	return db, nil
}

// Helper function to load TLS certificate
func mustLoadCert(certPath string) *x509.CertPool {
	certPool := x509.NewCertPool()
	cert, err := os.ReadFile(certPath)
	if err != nil {
		fmt.Println("Error loading TLS certificate: ", err)
		panic(err)
	}
	if !certPool.AppendCertsFromPEM(cert) {
		fmt.Println("Failed to append certificate")
		panic("Invalid certificate")
	}
	return certPool
}
