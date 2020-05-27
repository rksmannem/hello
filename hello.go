package main


import (
	"fmt"
	"os/user"
)


func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %v\n", usr.Name)
}
