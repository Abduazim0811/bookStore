package main

import (
	admin "BookStore/internal/Admin"
	signin "BookStore/internal/Signin"
	signup "BookStore/internal/Signup"
	st "BookStore/internal/Store"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		num  uint
		num2 uint
	)
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Abdu0811 dbname=users sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("[1]Users\t[2]Admin\t[3]Exit")
	fmt.Scanln(&num)

	switch num {
	case 1:
		fmt.Println("[1]SignIn   [2]Signup")
		fmt.Scanln(&num2)
		if num2 == 1 {
			signin.SignIn(db)
			if signin.Lampochka{
				st.Store(db)
			}
		} else if num2 == 2 {
			signup.SignUp(db)
		} else {
			fmt.Println("Notugri raqam kiritdiz!!!")
		}
	case 2:
		admin.Admin()
	case 3:
		os.Exit(0)
	}
}
