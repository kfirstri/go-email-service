package main

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/kfirstri/go-email-service/models"
)

func main() {
	fmt.Println("Hello!")
	u := make(map[int]models.User)

	var us models.User

	us.UserID = 1
	us.Firstname = "Kfir"
	us.Lastname = "Stri"
	us.Email = "kfir@gmail.com"
	us.Company = "world"

	u[1] = us

	fmt.Println(u[1])

	str := "1,Michael,Gelfand,michael@senexx.com,Senexx"

	r := csv.NewReader(strings.NewReader(str))

	record, err := r.Read()

	if err != nil {
		fmt.Println(err)
	}

	us, err = models.MakeUserFromRecord(record)

	fmt.Println(us)
}
