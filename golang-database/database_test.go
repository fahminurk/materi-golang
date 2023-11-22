package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectDb(t *testing.T) {
	db,err := sql.Open("mysql", "root:password@tcp(localhost:3306)/db_go")

	if err != nil {
		panic(err)
	}

defer db.Close()
}

/*
SetMaxIdleConns(number) = jumlah koneksi minimal yg dibuat
SetMaxOpenConns(number) = maksimal koneksi yg dibuat
SetConnMaxLifetime(number) = berapa lama koneksi boleh di jalankan
SetConnMaxIdleTime(number) = berapa lama koneksi yg sudah tidak digunakan akan dihapus
*/