package user

import "crud-web-server/db"

func InsertUser(user User) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", user.ID, user.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
