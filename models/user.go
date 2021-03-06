package models

import (
	"github.com/th3-z/malgo/storage"
)

type User struct {
	Id      int64
	Name    string
	Reviews []*Review
}

func NewUser(db storage.Queryer, name string) *User {
	query := `
        INSERT INTO user (name)
        VALUES (?)
    `

	userId, err := storage.PreparedExec(db, query, name)
	if err != nil {
		return SearchUser(db, name)
	}
	return GetUser(db, userId)
}

func GetUser(db storage.Queryer, userId int64) *User {
	query := `
        SELECT
            user_id,
            name
        FROM
            user
        WHERE
            user_id = ?
    `

	row := storage.PreparedQueryRow(
		db, query, userId,
	)
	var user User
	row.Scan(
		&user.Id, &user.Name,
	)

	user.Reviews = getUserReviews(db, userId)

	return &user
}

func SearchUser(db storage.Queryer, name string) *User {
	query := `
        SELECT
            user_id,
            name
        FROM
            user
        WHERE
            name = ?
    `

	row := storage.PreparedQueryRow(
		db, query, name,
	)
	var user User
	row.Scan(
		&user.Id, &user.Name,
	)

	user.Reviews = getUserReviews(db, user.Id)

	return &user
}

func (user *User) Update(db storage.Queryer) {
	query := `
        UPDATE user SET
			name = ?
        WHERE
			user_id = ?
    `

	_, err := storage.PreparedExec(db, query, user.Name, user.Id)
	if err != nil {
		panic(err)
	}

	for _, review := range user.Reviews {
		review.Update(db)
	}
}

func (user *User) Delete(db storage.Queryer) {
	query := `
        DELETE FROM user
        WHERE
			user_id = ?
    `

	_, err := storage.PreparedExec(db, query, user.Id)
	if err != nil {
		panic(err)
	}
}
