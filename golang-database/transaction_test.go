package golangdatabase

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransaction(t *testing.T){
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	tx,err := db.Begin()
	script := "INSERT INTO comments(email, comment) values(?, ?)"

	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		email := "admin" + strconv.Itoa(i) + "@gmail.com"
		comment := "komen ke " + strconv.Itoa(i)

	result, err :=	tx.ExecContext(ctx,script, email, comment)
		
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("comment id", id)
	}

	err = tx.Commit()

	if err != nil {
		panic(err)
		
	}

	tx.Rollback()
}