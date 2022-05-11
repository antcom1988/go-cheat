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

var sQLOpen = sql.Open

func OpenDBMonkeyPatch(user, password, addr, db string) (*sql.DB, error){
	conn := fmt.Sprintf("%s:%s@%s/%s", user, password, addr, db)
	return sQLOpen("mysql", conn)
}


/*
Source : https://www.youtube.com/watch?v=LEnXBueFBzk

Good
- Don't need to modify function signature

Not so good
- Allergic to parallelism (stateful)
- Have to make variable public (if testing form _test package (different directory))
*/