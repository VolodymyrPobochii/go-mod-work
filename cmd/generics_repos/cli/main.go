package main

import (
	"fmt"
	"github.com/VolodymyrPobochii/go-mod-work/internal"
)

func main() {
	usersRepo := internal.NewUsersRepo()

	u, err := usersRepo.FindByEmail("fgadgadf@davadfv.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(*u)
}
