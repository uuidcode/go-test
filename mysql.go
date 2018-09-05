package main

/**
 * https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.2.html
 */
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uuidcode/coreutil"
)

func main() {
	db, err := sql.Open("mysql", "root:rootroot@tcp(127.0.0.1:3306)/querydsl?charset=utf8")
	coreutil.CheckErr(err)

	sql := `
	insert into book 
	(mod_datetime, name, reg_datetime, user_id) 
	values (now(), ?, now(), ?)
	`

	stmt, err := db.Prepare(sql)
	coreutil.CheckErr(err)

	name := coreutil.CreateUuid()
	coreutil.CheckErr(err)

	res, err := stmt.Exec(name, 1)
	coreutil.CheckErr(err)

	affected, err := res.RowsAffected()
	coreutil.CheckErr(err)

	fmt.Println("affected", affected)

	stmt.Close()
	db.Close()
}
