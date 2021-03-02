package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("APP_DB_USERNAME", "postgres")
	os.Setenv("APP_DB_PASSWORD", "root")
	os.Setenv("APP_DB_NAME", "postgres")

	fmt.Println("APP_DB_USERNAME:", os.Getenv("APP_DB_USERNAME"))
	fmt.Println("APP_DB_PASSWORD:", os.Getenv("APP_DB_PASSWORD"))
	fmt.Println("APP_DB_NAME:", os.Getenv("APP_DB_NAME"))

	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
