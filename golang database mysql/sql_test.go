package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(id, name) VALUES ('aca', 'Aca')"
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name FROM CUSTOMER"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name : ", name)
	}

}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at  FROM CUSTOMER"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("============")
		fmt.Println("Id :", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		if birth_date.Valid {
			fmt.Println("Birth Date : ", birth_date.Time)
		}

		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", created_at)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// contoh sql injection

	// username := "admin'; #"
	// password := "slh"

	username := "admin"
	password := "admin"
	script := "SELECT username FROM user where username = '" + username + "' and password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("success login", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// contoh sql injection

	// username := "admin'; #"
	// password := "slh"

	username := "admin"
	password := "admin"
	script := "SELECT username FROM user where username = ? and password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}

		fmt.Println("success login", username)
	} else {
		fmt.Println("Gagal login")
	}

}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin2"
	password := "admin2"
	script := "INSERT INTO user(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "admin@gmail.com1"
	comment := "hallo"
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new comment", insertId)
}

// function untuk lgsg insert data sekaligus
func TestPrepareStmt(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// email := "admin@gmail.com1"
	// comment := "hallo"
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "admin" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini komen ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes insert new commentid : ", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// email := "admin@gmail.com1"
	// comment := "hallo"

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	for i := 0; i < 10; i++ {
		email := "admin" + strconv.Itoa(i) + "@gmail.com"
		comment := "ini komen ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Succes insert new commentid : ", insertId)
	}

	// rollback untuk membatalkan transaction
	// err = tx.Rollback()
	// if err != nil {
	// 	panic(err)
	// }
	// commit untuk melakukan transaction
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
