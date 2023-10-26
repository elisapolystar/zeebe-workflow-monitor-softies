package main


import (
	"fmt"
	"database/sql"
	"log"
	"github.com/lib/pq"
)

const (
	host = "postgres"
	port = "5432"
	user = "postgres"
	password = "password"
	dbname = "postgres"
	new_db = "kafka"
	create_table = "CREATE TABLE franz ( id integer, nimi varchar(255), sivumaara integer, vuosi integer);"
	insertion_1 = "INSERT INTO franz (id, nimi, sivumaara, vuosi) VALUES (1, 'Anna Karenina', 864, 1878),
	(2, 'Madame Bovary', 322, 1856), (3, 'War and Peace', 1225, 1869), (4, 'The Great Gatsby', 218, 1925),
	(5, 'Lolita', 371, 1959), (6, 'Middlemarch', 501, 1872), (7, 'The Adventures of Huckleberry Finn', 362, 1884),
	(8, 'The Stories of Antony Chekhov', 269, 1885), (9, 'In Search of Lost Time', 4215, 1913), (10, 'Hamlet', 500, 1600);"
	display_1 = "FROM franz SELECT *;"
	deletion = "DELETE FROM 'franz' WHERE sivumaara > 500;"
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

	// Create a database named "kafka"
	_, err = db.Exec("create database " + new_db)
	if err != nil {
		fmt.Println("Database creation failed")
		panic(err)
	}

	// Create a table into Kafka named "franz"
	_, err = db.Exec(create_table)
	if err != nil {
		fmt.Println("Table creation failed")
		panic(err)
	}
	// insert 10 entries into the table
	_, err = db.Exec(insertion_1)
	if err != nil {
		fmt.Println("Data insertion into database failed")
		panic(err)
	}
	// display the current table
	rows, err = db.Exec(display_1)
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
	_, err = db.Exec(deletion);
	if err != nil {
		fmt.Println("Failed to delete rows")
		panic(err)
	}
	// display the modified table
	rows, err = db.Exec(display_1)
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

}


