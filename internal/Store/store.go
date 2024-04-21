package store

import (
	// us "BookStore/models"
	"database/sql"
	"fmt"
	"log"
	"os"
)


func Store(db *sql.DB, email string) {
	fmt.Println("Kitoblar ro'yxati:")

	rows, err := db.Query("SELECT id, name, author, count, price FROM book WHERE count > 0")
	if err != nil {
		log.Fatalf("Kitoblarni olishda xato: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, count int
		var name, author string
		var price float64
		if err := rows.Scan(&id, &name, &author, &count, &price); err != nil {
			log.Fatalf("Kitob ma'lumotlarini o'qishda xato: %v", err)
		}
		fmt.Printf("%d. %s - %s, qoldiq: %d, narxi: %.2f\n", id, name, author, count, price)
	}

	var num int
	fmt.Print("\nTanlangan kitob raqamini kiriting: ")
	fmt.Scanln(&num)


	var bookID int
	var name, author string
	var count int
	var bookPrice float64
	err = db.QueryRow("SELECT id, name, author, count, price FROM book WHERE id = $1", num).Scan(&bookID, &name, &author, &count, &bookPrice)
	if err != nil {
		log.Fatalf("Kitob topilmadi: %v", err)
	}

	fmt.Printf("\nTanlangan kitob: %s - %s, narxi: %.2f\n", name, author, bookPrice)

	var balance float64
	err = db.QueryRow("SELECT price FROM users WHERE email = $1", email).Scan(&balance)
	// fmt.Println(user.Email)
	if err != nil {
		log.Fatalf("Foydalanuvchi balansini olishda xato: %v", err)
	}

	if bookPrice <= balance {
		_, err = db.Exec("UPDATE book SET count = count - 1 WHERE id = $1 AND count > 0", bookID)
		if err != nil {
			log.Fatalf("Kitob zaxirasini yangilashda xato: %v", err)
		}
		_, err = db.Exec("UPDATE users SET price = price - $1 WHERE email = $2", bookPrice, email)
		if err != nil {
			log.Fatalf("Foydalanuvchi balansini yangilashda xato: %v", err)
		}
		fmt.Println("Kitob sotildi")
	} else {
		fmt.Println("\nHisobizda mablag' yetarli emas!! Boshqa kitoblarni kurishni hohlaysizmi?")
		fmt.Println("[1] HA   [2]Yoq")
		fmt.Scanln(&num)
		if num == 1 {
			Store(db,email) 
		} else {
			os.Exit(0)
		}
	}
}
