package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T){
	db := GetConnection()
	defer db.Close()
	
	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('2', 'budi')"

	_,err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")

}

func TestQueryContex(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, created_at, birth_date, married FROM customer"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name  string
		var email sql.NullString // jika data berupa null
		var balance sql.NullInt32 // int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, " name:", name, " email:", email, " balance:", balance, " rating:", rating, " created_at:", created_at, " birth_date:", birth_date, " married:", married)
	}

	defer rows.Close()
}

func TestSqlInjectionSafe(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "admin" // input dari user

	query := "SELECT username from user where username = ? and password = ? limit 1" // ganti dengan tanda ?

	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
		
}
func TestSqlInjection(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah" // input dari user

	query := "SELECT username from user where username = '" + username + "' and password = '" + password + "' limit 1"

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login", username)
	} else {
		fmt.Println("gagal login")
	}
		
}
func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "admin"
	comment := "komen" // input dari user

	query := "insert into comments(email, comment) values(?, ?)"

	result, err := db.ExecContext(ctx, query, email, comment)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
		
	fmt.Println("Success insert new comment with id", id)
}

func TestPrepareStmt(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	statement, err := db.PrepareContext(ctx, "INSERT INTO comments(email, comment) values(?, ?)")

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "admin" + strconv.Itoa(i) + "@gmail.com"
		comment := "komen ke " + strconv.Itoa(i)

	result, err :=	statement.ExecContext(ctx, email, comment)
		
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("comment id", id)
	}

}