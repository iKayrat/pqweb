package models

import (
	_ "github.com/lib/pq"
)

// func init() {
// 	orm.RegisterModel(new(Cidade))
// }

// func NewDBconnection() (con *pg.DB) {
// 	address := fmt.Sprintf("%s:%s", "localhost", "5432")
// 	options := &pg.Options{
// 		User:     "postgres",
// 		Password: "kaak",
// 		Addr:     address,
// 		Database: "firstapp",
// 	}
// 	con = pg.Connect(options)
// 	if con == nil {
// 		panic("Cannot connect to postgres")
// 	}

// 	return
// }

// // DBConnection ...
// type DBConnection struct {
// 	DB *sql.DB
// }

// var Profile *parent.Profile

// // Connection ...
// var Connection DBConnection

// // ConnectionDB ...
// func ConnectionDB() {
// 	db, err := sql.Open("postgres", "postgres://postgres://postgres:kaak@localhost/firstapp?sslmode=disable")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connected to database")

// 	defer db.Close()
// }
