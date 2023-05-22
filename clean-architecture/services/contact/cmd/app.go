package main

import (
	"fmt"

	"alexxx.go-cleanArch/pkg/store/postgres"
	"alexxx.go-cleanArch/services/contact/internal/domain"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "1234567",
		DbName:   "Day1",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	alexey := domain.NewContact("Fedenko", "Alexey", "Demyanovich")
	nurdaulet := domain.NewContact("Nurdaulet", "Kuatov", "Ivanovich")
	group1 := domain.NewGroup("Students")

	fmt.Println(alexey, nurdaulet, group1)
}
