package main

import (
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

func main() {
	hash := userfields.Password(`$2a$10$7sd/YfNSXu1G0l5vYAOkGumMR3Jsw.tS9ZJIY.mBgnNPQvEpkxAka`)
	password := "voronkovda887"
	fmt.Println(hash.Verify(password))
}
