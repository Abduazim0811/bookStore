package signin

import (
	us "BookStore/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func Scan(user *us.Users) {
	fmt.Printf("Emal: ")
	fmt.Scanln(&user.Email)
	fmt.Printf("Password: ")
	fmt.Scanln(&user.Password)
}

var Lampochka bool

func SignIn(db *sql.DB) {
	var user us.Users
	Scan(&user)

	query := "SELECT (email,password) FROM users;"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var email, password string
		if err := rows.Scan(&email, &password); err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(user.Email) == strings.TrimSpace(email) && strings.TrimSpace(user.Password) == strings.TrimSpace(password) {
			Lampochka = true
			return
		}
	}

}
