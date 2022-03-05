package db

import (
	"database/sql"
	"fmt"
)

type User struct {
	Login    string
	Password string
	Name     string
}

var query map[string]*sql.Stmt

func prepare() error {
	query = make(map[string]*sql.Stmt)

	var e error
	query["UserInsert"], e = Link.Prepare(`INSERT INTO "User"
		("Login", "Password", "Name")
		VALUES ($1, $2, $3)`)
	if e != nil {
		fmt.Println(e)
		return e
	}

	query["UserSelect"], e = Link.Prepare(
		`SELECT "Login", "Password", "Name" FROM "User" ORDER BY "Name"`)
	if e != nil {
		fmt.Println(e)
		return e
	}

	return nil
}

func (u *User) Insert() {
	tx, e := Link.Begin()
	if e != nil {
		return
	}

	defer deferTx(tx)

	stmt, ok := query["UserInsert"]
	if !ok {
		return
	}
	_, e = stmt.Exec(u.Login, u.Password, u.Name)
	if e != nil {
		Logger.Println(e)
		panic(e)
	}
}

func (u *User) Select() {
	row := Link.QueryRow(`SELECT "Name" FROM "User" 
WHERE "Login"=$1 AND "Password"=$2`,
		u.Login, u.Password)
	e := row.Scan(&u.Name)
	if e != nil {
		Logger.Println(e)
	}
}

func (u *User) SelectAll() []User {
	rows, e := Link.Query(`SELECT "Login", "Name" FROM "User" ORDER BY "Name"`)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		e = rows.Scan(&u.Login, &u.Name)
		if e != nil {
			fmt.Println(e)
			return nil
		}

		users = append(users, User{
			Login:    u.Login,
			Password: "",
			Name:     u.Name,
		})
	}

	return users
}

func deferTx(tx *sql.Tx) {
	r := recover()
	if r != nil {
		_ = tx.Rollback()
	} else {
		_ = tx.Commit()
	}
}
