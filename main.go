package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/CharVstack/ChaV-lib/host"
)

func init() {
	godotenv.Load("./.env")
}

func main() {
	hoge := host.GetInfo()
	fmt.Printf("%+v", hoge)
}
