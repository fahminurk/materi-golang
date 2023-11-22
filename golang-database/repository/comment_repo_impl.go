package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB //pointer
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {

result,err := repository.DB.ExecContext(ctx, "INSERT INTO comments(email, comment) values(?, ?)", comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

id, err := result.LastInsertId()
comment.Id = int32(id)

return comment, nil
}
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {

	rows, err := repository.DB.QueryContext(ctx, "SELECT id, email, comment FROM comments WHERE id = ? limit 1", id)
	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("id " + strconv.Itoa(int(id)) + " not found")
	}
}
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {

	rows, err := repository.DB.QueryContext(ctx, "SELECT id, email, comment FROM comments")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var comments []entity.Comment

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
	
}