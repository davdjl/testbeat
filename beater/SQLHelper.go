package beater


import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "192.168.1.64"
var port = 1433
var user = "golang"
var password = "jackcloudman"
var database = "Test"

// Person is a date type that contains people info
type Person struct {
	name     string
	location string
	id       int
}

func startConection() {
	// INICIAMOS CONEXION
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
}

// ReadEmployees reads all employee records
func ReadEmployees(currentID int) ([]Person, error) {
	ctx := context.Background()

	// Check if database is alive.
	err := db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees WHERE id>@ID;")

	// Execute query
	rows, err := db.QueryContext(ctx, tsql, sql.Named("ID", currentID))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var persons []Person
	// Iterate through the result set.
	for rows.Next() {
		var name, location string
		var id int

		// Get values from row.
		err := rows.Scan(&id, &name, &location)
		p := Person{name: name, id: id, location: location}
		persons = append(persons, p)

		if err != nil {
			return nil, err
		}

		//fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
	}
	return persons, nil
}
