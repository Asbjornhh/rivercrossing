package main

import (
	"fmt"
	"github.com/Asbjornhh/rivercrossing/event"
	"github.com/Asbjornhh/rivercrossing/state"
)

func main() {
	fmt.Println(state.ViewState())
	fmt.Println(event.PutIn())
	fmt.Println(event.CrossRiver())
	fmt.Println(event.TakeOut())
}
