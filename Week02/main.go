package main

import (
	"database/sql"
	"errors"
)

type User struct {
	Id   int
	Name string
}

// dao: wrap errors
func getId(id int) (*User, error) {
	db, err := sql.Open("mysql", "kuma:password@tcp(127.0.0.1:3306)/goku")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer db.Close()
	user := &User{}
	err = db.QueryRow("select id, name from users where id = ? ", id).Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.WithStack(err)
	}
	return user, nil
}

// controller: use errors
func getUserById(ident int) (*User, error) {
	return getId(ident)
}


// api: get errors status
func queryUser(id int) (int, *User) {
	user, err := getUserById(id)
	if err != nil {
		// logger.Error("api_fail, err="+err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			return 404, user
		}
		return 500, user
	}
	return 200, user
}
