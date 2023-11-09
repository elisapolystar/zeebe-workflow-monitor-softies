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
	//check what type of entity we are saving
	switch d := entity.(type) {
	//save a process entity
	case Process:
		process := d;
		fmt.Println("saving process")
		insertProcess := `INSERT INTO process (Key, BpmnProcessId, Version, Resource, Timestamp) VALUES ($1, $2, $3, $4, $5)`
		//execute the insertion command with entity as parameters
		_, err = db.Exec(insertProcess, process.Key, process.Value.BpmnProcessId, process.Value.Version, process.Value.Resource, process.Timestamp)
		if err != nil {
			fmt.Println("Data insertion into database failed")
			fmt.Println(err)
		} else {
			fmt.Println("saved process to database!")
		}
	//save an instance entity
	case ProcessInstance:
		fmt.Println("saving instance")
		//TODO add statement
	default:
        fmt.Println("Unsupported entity")
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
	host, port, user, password, DBname)

	// Open a database connection, and check that it works
	db, err := sql.Open("postgres", DBConnection)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database!")
	return db, nil
}

//manual table creation REDUNDANT!!!
func CreateTables() {
	db, err := connectToDatabase()
    if err != nil {
        fmt.Println("Error opening database connection:", err)
    }
	create_process := "CREATE TABLE process ( Key BIGINT, BpmnProcessId VARCHAR(50) NOT NULL, Version INT NOT NULL, Resource TEXT NOT NULL, Timestamp BIGINT NOT NULL);"
	_, err = db.Exec(create_process)
	if err != nil {
		fmt.Println("Table creation failed")
	}
}