package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepo := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "fahmi@mail.com",
		Comment: "komen",
	}
	
	result, err := commentRepo.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}


func TestFindById(t *testing.T){
	commentRepo := NewCommentRepository(golangdatabase.GetConnection())
	ctx := context.Background()

	comment, err := commentRepo.FindById(ctx, 99)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}