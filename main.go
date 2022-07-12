package main

import (
	"github.com/joho/godotenv"

	_ "github.com/CharVstack/CharV-lib/host"
	_ "github.com/CharVstack/CharV-lib/qemu"
)

func init() {
	godotenv.Load("./.env")
}

func main() {
}
