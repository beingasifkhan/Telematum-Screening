package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"screening/database"
)

func SetupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn, err := database.CreateConnection()
		if err != nil {
			log.Println("error establishing database connection:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// /createUser
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "INSERT INTO users (name, email) VALUES ($1, $2)"
		result, err := conn.Exec(query, name, email)
		if err != nil {
			log.Println("error executing createUser query:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("result:", result)
		w.Write([]byte("Created user successfully!"))

	})
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn, err := database.CreateConnection()
		if err != nil {
			log.Println("error establishing database connection:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// /updateUser
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("id")

		// Check if the user with the specified ID exists
		exists, err := userExists(conn, id)
		if err != nil {
			log.Println("error checking if user exists:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// User exists, proceed with the update
		query := "UPDATE users SET name=$1, email=$2 WHERE id=$3"
		result, err := conn.Exec(query, name, email, id)
		if err != nil {
			log.Println("error executing updateUser query:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("result:", result)
		w.Write([]byte("User updated successfully!"))

	})

}
func userExists(conn *sql.DB, userID string) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE id = $1"
	err := conn.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
