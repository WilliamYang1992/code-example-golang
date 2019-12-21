package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Host     = "127.0.0.1"
	Port     = 3306
	User     = "root"
	Password = "password"
	DB       = "world"
)

func Example() {
	driver := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		User, Password, Host, Port, DB)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	stmt := `SELECT code, name, continent, region, indepYear 
			 FROM country WHERE NAME = ?`
	rows, err := db.Query(stmt, "China")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		code             string
		name             string
		continent        string
		region           string
		independenceYear sql.NullInt32
	)
	for rows.Next() {
		if err := rows.Scan(&code, &name, &continent, &region,
			&independenceYear); err != nil {
			panic(err)
		}
		fmt.Printf("code: %s, name: %s\n", code, name)
	}
	// Output:
	// code: CHN, name: China
}
