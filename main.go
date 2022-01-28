package main

import (
	"fmt"
)

func main() {
	fmt.Println(state.ViewState())
	fmt.Println(event.PutIn("Kylling"))
	fmt.Println(event.CrossRiver("Kylling"))
	fmt.Println(event.TakeOut("Kylling"))
}
