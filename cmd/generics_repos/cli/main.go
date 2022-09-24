package main

import (
	"fmt"
	"github.com/VolodymyrPobochii/go-mod-work/cmd/generics_repos"
	"google.golang.org/grpc/status"
)

func main() {
	usersRepo := generics_repos.NewUsersRepo()

	u, err := usersRepo.FindByEmail("fgadgadf@davadfv.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(*u)

	status.Code()
}
