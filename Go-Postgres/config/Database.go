// Database Configuration
package Database

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"  

)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0719076633"
	dbname   = "testdatabase"
  )

func CreateConnection() *sql.DB {
	
   psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db,err := sql.Open("postgres",psqlInfo)

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}
  


