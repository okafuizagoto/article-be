// @title Example API
// @version 1.0
// @description goldgym Example API
// @BasePath /example
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"article-be/internal/boot"
	"article-be/internal/config"
	articles "article-be/internal/entity/article"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := boot.HTTP(); err != nil {
		log.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
	dsn := "user:admin@tcp(localhost:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	// AutoMigrate akan membuat table sesuai struct jika belum ada
	if err := db.AutoMigrate(&articles.Post{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	log.Println("migration done")
}

func openDatabases(cfg *config.Config) (master *sqlx.DB, err error) {
	master, err = openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		return master, err
	}

	return master, err
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}
