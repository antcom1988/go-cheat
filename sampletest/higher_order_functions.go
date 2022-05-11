package sampletest

import (
	"database/sql"
	"fmt"
)

/* -> Before
func main(){
	db, _ := OpenDB("user", "password", "localhost:3304", "db")
}

func OpenDB(user, password, addr, db string) (*sql.DB, error){
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return sql.Open("mysql", conn)
}
*/

//-> After

type sqlOpener func (string,string) (*sql.DB, error)

func OpenDBHighOrderFunc(user, password, addr, db string, opener sqlOpener) (*sql.DB, error){
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return opener("mysql", conn)
}


/*
Source : https://www.youtube.com/watch?v=LEnXBueFBzk

Good
- Clear and proximal to function under test
- Stateless

Not so good
- Parameter list can get ugly(using a "type" for the function can help)
- Think of dependency graph
*/