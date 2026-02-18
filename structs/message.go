package structs

// import (
// 	"database/sql"
// 	"log"
// )

// type Message struct {
// 	Name    string
// 	Email   string
// 	Message string
// }

// var dbInstance *sql.DB

// func InitDb() error {
// 	var err error
// 	dbInstance, err = sql.Open("sqlite3", "./messages.db")

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func GetDb() *sql.DB {
// 	return dbInstance
// }

// func CreateMessagesTable() {
// 	query := `
// 		CREATE TABLE IF NOT EXISTS messages (
// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
// 			name TEXT,
// 			email TEXT,
// 			message TEXT
// 		);
// 	`

// 	_, err := dbInstance.Exec(query)

// 	if err != nil {
// 		log.Fatalf("Error al crear la tabla: %v\n", err)
// 	}
// }

// func (m *Message) InsertMessage() error {
// 	stmt, err := dbInstance.Prepare("INSERT INTO messages (name, email, message) VALUES (?, ?, ?)")

// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(m.Name, m.Email, m.Message)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
