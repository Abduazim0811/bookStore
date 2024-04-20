package admin

import (
	book "BookStore/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func Scan(bk *book.Book) {
	fmt.Printf("Name:  ")
	fmt.Scanln(&bk.Name)
	fmt.Printf("Author:  ")
	fmt.Scanln(&bk.Author)
	fmt.Printf("Count:  ")
	fmt.Scanln(&bk.Count)
	fmt.Printf("Price:  ")
	fmt.Scanln(&bk.Price)
}

func Admin(db *sql.DB) {
	var (
		bk book.Book
		lampochka bool
	)
	query := "SELECCT *FROM book;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	Scan(&bk)

	for rows.Next() {
		var (
			id, count    int
			name, author string
			price        float64
		)
		if err := rows.Scan(&id, &name, &author, &count, &price); err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(name)==strings.TrimSpace(bk.Name)&&strings.TrimSpace(author)==strings.TrimSpace(bk.Author){
			lampochka=true
		}
	}

	if lampochka{
		_,err:=db.Query("UPDATE book SET count=count+$1 WHERE name=$2", bk.Count,bk.Name)
		if err!=nil{
			log.Fatal(err)
		}
	}else{
		_,err:=db.Query("INSERT INTO book(name,author,count,price) VALUES ($1,$2,$3,$4)",bk.Name,bk.Author,bk.Count,bk.Price)
		if err!=nil{
			log.Fatal(err)
		}
	}
}
