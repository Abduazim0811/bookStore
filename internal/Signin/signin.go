package signin

import (
	us "BookStore/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func Scan(user *us.Users){
	fmt.Printf("Email: ")
	fmt.Scanln(&user.Email)
	fmt.Printf("Password: ")
	fmt.Scanln(&user.Password)
}

var (
	Lampochka bool
	Emaill string
) 


func SignIn(db *sql.DB) {
	var (
		user  us.Users
		son int
	)
	Scan(&user)
	Emaill=user.Email
	query := "SELECT email,password,price FROM users;"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			email, password string
			price float64
		)
		if err := rows.Scan(&email, &password); err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(user.Email) == strings.TrimSpace(email) && strings.TrimSpace(user.Password) == strings.TrimSpace(password) {
			Lampochka = true
			fmt.Println("Hisobizda",price,"pul bor")
			fmt.Println("Hisobizni toldirishni hohlaysizmi?\n [1] Ha [2]Yoq")
			fmt.Scanln(&son)
			if son==1{
				fmt.Println("Pul kiriting: ?$")
			fmt.Scanln(&user.Pricee)
			_,err:=db.Query("UPDATE users SET price=price+$1 WHERE email=$2", price,user.Email)
			if err!=nil{
				log.Fatal(err)
			}
			}
		}else{
			fmt.Println("Bunday foydalanuvchi yoq!!")
		}
	}
}
