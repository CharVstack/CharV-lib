package main

import (
	"fmt"

	"github.com/CharVstack/ChaV-lib/host"
)

func main() {
	hoge := host.GetInfo()
	fmt.Printf("%+v", hoge)
}
