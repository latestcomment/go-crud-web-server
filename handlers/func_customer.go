package handlers

import (
	"crud-web-server/db"
	"crud-web-server/models"
	"crud-web-server/utils"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func getAllCustomers() ([]models.Customer, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	rows, err := tx.Query("SELECT customeruuid, firstname, middleinitial, lastname FROM customers LIMIT 5")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var tempCustomer struct {
			CustomerId    uuid.UUID
			FirstName     sql.NullString
			MiddleInitial sql.NullString
			LastName      sql.NullString
		}

		if scanErr := rows.Scan(&tempCustomer.CustomerId, &tempCustomer.FirstName, &tempCustomer.MiddleInitial, &tempCustomer.LastName); scanErr != nil {
			err = scanErr
			return nil, err
		}
		customers = append(customers, models.Customer{
			CustomerId:    tempCustomer.CustomerId,
			FirstName:     utils.NullToString(tempCustomer.FirstName),
			MiddleInitial: utils.NullToString(tempCustomer.MiddleInitial),
			LastName:      utils.NullToString(tempCustomer.LastName),
		})
	}

	return customers, nil
}

func getCustomer(id uuid.UUID) (*models.Customer, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	row := tx.QueryRow("SELECT customeruuid, firstname, middleinitial, lastname FROM customers WHERE customeruuid = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var tempCustomer struct {
		CustomerId    uuid.UUID
		FirstName     sql.NullString
		MiddleInitial sql.NullString
		LastName      sql.NullString
	}
	if err := row.Scan(&tempCustomer.CustomerId, &tempCustomer.FirstName, &tempCustomer.MiddleInitial, &tempCustomer.LastName); err != nil {
		return nil, err
	}
	customer := &models.Customer{
		CustomerId:    tempCustomer.CustomerId,
		FirstName:     utils.NullToString(tempCustomer.FirstName),
		MiddleInitial: utils.NullToString(tempCustomer.MiddleInitial),
		LastName:      utils.NullToString(tempCustomer.LastName),
	}

	return customer, nil
}

func updateCustomers(updates []map[string]interface{}) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	for _, update := range updates {
		id, ok := update["customeruuid"].(string)
		if !ok {
			return fmt.Errorf("customeruuid is required in update payload")
		}

		setClauses := []string{}
		args := []interface{}{}

		i := 1
		for col, val := range update {
			if col == "customeruuid" {
				continue
			}
			setClauses = append(setClauses, fmt.Sprintf(`"%s" = $%d`, col, i))
			args = append(args, val)
			i++
		}
		args = append(args, id)
		q := fmt.Sprintf(
			`UPDATE customers
             SET %s
             WHERE customeruuid = $%d`,
			strings.Join(setClauses, ", "),
			i,
		)
		if _, err = tx.Exec(q, args...); err != nil {
			return err
		}
	}
	return nil
}

func deleteCustomer(id uuid.UUID) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec("DELETE FROM customers WHERE customeruuid = $1", id)
	if err != nil {
		return err
	}
	return nil
}
