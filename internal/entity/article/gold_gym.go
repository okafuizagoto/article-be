package article

import (
	"time"
)

type Post struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`
	Content     string    `gorm:"type:text" json:"content"`
	Category    string    `gorm:"type:varchar(100)" json:"category"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);default:'Draft'" json:"status"`
}

type Put struct {
	Title    string `db:"title" json:"title"`
	Content  string `db:"content" json:"content"`
	Category string `db:"category" json:"category"`
	Status   string `db:"status" json:"status"`
}

type Get struct {
	ID          int    `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Content     string `db:"content" json:"content"`
	Category    string `db:"category" json:"category"`
	CreatedDate string `db:"created_date" json:"created_date"`
	UpdatedDate string `db:"updated_date" json:"updated_date"`
	Status      string `db:"status" json:"status"`
}

type GetGoldUser struct {
	GoldId            int    `db:"gold_id" json:"gold_id"`
	GoldEmail         string `db:"gold_email" json:"gold_email"`
	GoldPassword      string `db:"gold_password" json:"gold_password"`
	GoldNama          string `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp       string `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu    string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv           string `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate   string `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu string `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
}

type GetGoldUsers struct {
	GoldId            int    `db:"gold_id" json:"gold_id"`
	GoldEmail         string `db:"gold_email" json:"gold_email"`
	GoldPassword      string `db:"gold_password" json:"gold_password"`
	GoldNama          string `db:"gold_nama" json:"gold_nama"`
	GoldNomorHp       string `db:"gold_nomorhp" json:"gold_nomorhp"`
	GoldNomorKartu    string `db:"gold_nomorkartu" json:"gold_nomorkartu"`
	GoldCvv           string `db:"gold_cvv" json:"gold_cvv"`
	GoldExpireddate   string `db:"gold_expireddate" json:"gold_expireddate"`
	GoldPemegangKartu string `db:"gold_namapemegangkartu" json:"gold_namapemegangkartu"`
	GoldOTP           string `db:"gold_otp" json:"gold_otp"`
}
