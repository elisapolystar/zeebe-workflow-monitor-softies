package main


import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
	user = "postgres"
	password = "password"
	DBname = "workflow"
	
	ProcessesQuery = "SELECT * FROM process;"
)

func SaveData(entity interface{}) {
	//connect to database
	db, err := connectToDatabase()
    if err != nil {
        fmt.Println("Error opening database connection when saving:", err)
    }
	//check what we are saving
	switch entity.(type) {
		//save a process entity
		case Process:
			fmt.Println("saving process")
			//prepare the sql %d = number %s = text/string
			//insertProcess = fmt.Sprintf("INSERT INTO process (Key, BpmnProcessId, Version, Resource, timestamp) VALUES (%d, '%s', %d, '%s', %d);";
			insertProcess := fmt.Sprintf("INSERT INTO process (Key, BpmnProcessId, Version, Resource, Timestamp) VALUES ($1, $2, $3, $4, $5)")
			//execute the insertion command with entity as parameters
			_, err = db.Exec(insertProcess, entity.Key, entity.Value.BpmnProcessId, entity.Value.Version, entity.Value.Resource, entity.Timestamp)
			if err != nil {
				fmt.Println("Data insertion into database failed")
			} else {
				fmt.Println("saved process to database!")
			}
		//save an instance entity
		case Instance:
			fmt.Println("saving instance")
			//TODO add statement
	}
}

func RetrieveProcesses() *sql.Rows {
	fmt.Println("retrieving processes from the database")
	//connect to database
	db, err := connectToDatabase()
    if err != nil {
        fmt.Println("Error opening database connection:", err)
    }
	fmt.Println("processes retrieved succesfully")
	rows, err := db.Query(ProcessesQuery)
	defer rows.Close()
	return rows
}

func connectToDatabase() (*sql.DB, error){
	//pass variables to the connection string
	DBConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, password, dbname)

	// Open a database connection, and check that it works
	db, err := sql.Open("postgres", db_conn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	fmt.Println("Connected to the database!")
	return db, nil
}