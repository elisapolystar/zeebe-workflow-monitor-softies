package main


import (
	"fmt"
	"database/sql"
	//"log"
	_ "github.com/lib/pq"
)

const (
	host = "postgres"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "test"
	new_db = "kafka"
	
	create_table = "CREATE TABLE franz ( id integer, nimi varchar(255), sivumaara integer, vuosi integer);"
	insertion_1 = "INSERT INTO franz (id, nimi, sivumaara, vuosi) VALUES (1, 'Anna Karenina', 864, 1878)," +
	"(2, 'Madame Bovary', 322, 1856), (3, 'War and Peace', 1225, 1869), (4, 'The Great Gatsby', 218, 1925)," +
	"(5, 'Lolita', 371, 1959), (6, 'Middlemarch', 501, 1872), (7, 'The Adventures of Huckleberry Finn', 362, 1884)," +
	"(8, 'The Stories of Antony Chekhov', 269, 1885), (9, 'In Search of Lost Time', 4215, 1913), (10, 'Hamlet', 500, 1600);"
	display_1 = "SELECT * FROM franz;"
	deletion = "DELETE FROM franz WHERE sivumaara > 500;"
	
)


func TestDatabase() {
	// Create a string to connect to the DB
	db_conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, password, dbname)

	// Open a database connection, and check that it works
	db, err := sql.Open("postgres", db_conn)
	if err != nil {
		panic(err)
	  }
	defer db.Close()
	fmt.Println("Connected to the database!")

	//Create a database named "kafka"
	fmt.Println("creating new database")
	createDB := "CREATE DATABASE " + new_db
	_, err = db.Exec(createDB)
	if err != nil {
		fmt.Println("Database creation failed")
		panic(err)
	}
	// connect to the new database
	new_db_conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, password, new_db)
	db_2, err2 := sql.Open("postgres", new_db_conn)
	if err2 != nil {
		fmt.Println("connecting to the new database failed")
		panic(err)
	}
	// Create a table into Kafka named "franz"
	fmt.Println("creating table")
	_, err = db.Exec(create_table)
	if err != nil {
		fmt.Println("Table creation failed")
		panic(err)
	}
	// insert 10 entries into the table
	fmt.Println("Inserting rows...")
	_, err = db.Exec(insertion_1)
	if err != nil {
		fmt.Println("Data insertion into database failed")
		panic(err)
	}
	// display the current table
	fmt.Println("displaying results")
	rows, err := db.Query(display_1)
	defer rows.Close()
	for rows.Next(){
		var id int
		var nimi string
		var sivumaara int
		var vuosi int

		err = rows.Scan(&id, &nimi, &sivumaara, &vuosi)
		if err != nil {
			fmt.Println("Failed to scan a row")
			panic(err)
		}
		fmt.Println(id, nimi, sivumaara, vuosi)
	}
	// delete every book containing more than 500 pages from the table
	fmt.Println("deleting too long books...")
	_, err = db.Exec(deletion)
	if err != nil {
		fmt.Println("Failed to delete rows")
		panic(err)
	}
	// display the modified table
	fmt.Println("displaying results...")
	rows, err = db.Query(display_1)
	defer rows.Close()
	for rows.Next(){
		var id int
		var nimi string
		var sivumaara int
		var vuosi int

		err = rows.Scan(&id, &nimi, &sivumaara, &vuosi)
		if err != nil {
			fmt.Println("Failed to scan a row")
			panic(err)
		}
		fmt.Println(id, nimi, sivumaara, vuosi)
	}
	fmt.Println("dropping database...")
	_, err = db_2.Exec("DROP DATABASE kafka")
	if err != nil {
		fmt.Println("Failed to drop the database")
		panic(err)
	}
	fmt.Println("Database Dropped")
	
}


