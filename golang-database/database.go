package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB{
	db,err := sql.Open("mysql", "root:password@tcp(localhost:3306)/db_go?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) // set jumlah koneksi awal yg standby
	db.SetMaxOpenConns(100) // set jumlah maximal koneksi yg dibuat
	db.SetConnMaxIdleTime(5 * time.Minute) // set berapa lama koneksi yg tidak digunakan akan di hapus
	db.SetConnMaxLifetime(60 * time.Minute) // set berapa lama koneksi boleh di jalankan

	return db
}