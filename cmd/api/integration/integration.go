package integration

import (
	"beto_integration/cmd/api/operations"
	"fmt"
)

func Run(){
	fmt.Println("Dude i'm running")
	operations.Purchase()
}