package store

import (
	"database/sql"
	"fmt"
	"log"
)

func Store(db *sql.DB) {
	var (
		num int
		id  int
	)
	query := "SELECT * FROM book;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Kitoblar ro'yxati:")
	for rows.Next() {
		var name, author string
		var count int
		var price float64
		err := rows.Scan(&id, &name, &author, &count, &price)
		if err != nil {
			log.Fatal(err)
		}
		if count!=0{
			fmt.Printf("%d. %s - %s, qoldiq: %d, narxi: %.2f\n", id, name, author, count, price)
		}
	}

	fmt.Print("Tanlangan kitob raqamini kiriting: ")
	fmt.Scanln(&num)
	var selectedID int
	rows, err = db.Query("SELECT id FROM book WHERE id = $1", num)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&selectedID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatalf("Raqam %d bo'yicha kitob topilmadi", num)
	}

	var (
		name   string
		author string
		count  int
		price  float64
	)
	err = db.QueryRow("SELECT name, author, count, price FROM book WHERE id = $1", selectedID).Scan(&name, &author, &count, &price)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tanlangan kitob: %s - %s, narxi: %.2f\n", name, author, price)

	_, err = db.Exec("UPDATE book SET count = count - 1 WHERE name = $1 AND count > 0", name)
	if err != nil {
		log.Fatal(err)
	}

}

