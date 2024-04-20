package signup

import (
	users "BookStore/models"
	"database/sql"
	"fmt"
	"log"
)

func Scan(us *users.Users) {
	fmt.Printf("First_name: ")
	fmt.Scanln(&us.First_Name)
	fmt.Printf("Last_name: ")
	fmt.Scanln(&us.Last_Name)
	fmt.Printf("Email: ")
	fmt.Scanln(&us.Email)
	fmt.Printf("Password: ")
	fmt.Scanln(&us.Password)
}

var Lampochka bool

func SignUp(db *sql.DB) {
	var (
		user users.Users
		price float64
	)

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users(
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(50),
			last_name VARCHAR(50),
			email VARCHAR(50) UNIQUE NOT NULL,
			password VARCHAR(50) NOT NULL
			price FLOAT NOT NULL
		)`)
	if err != nil {
		log.Fatal(err)
	}
	Scan(&user)

	query := "INSERT INTO users(first_name, last_name,email,password,price) VALUES($1,$2,$3,$4,$5)"
	_, err = db.Exec(query, user.First_Name, user.Last_Name, user.Email, user.Password, 0.0)
	if err != nil {
		log.Fatal(err)
	}
	Lampochka = true
	fmt.Println("Pul kiriting: ?$")
	fmt.Scanln(&price)
	_, err = db.Query("UPDATE users SET price=price+$1 WHERE email=$2", price, user.Email)
	if err != nil {
		log.Fatal(err)
	}

}
