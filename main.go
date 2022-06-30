package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/CharVstack/CharV-lib/host"
)

func init() {
	godotenv.Load("./.env")
}

func main() {
	hoge := host.GetInfo()
	fmt.Printf("%+v", hoge)
}
