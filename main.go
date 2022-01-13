package main

import (
	explorer "github.com/magutae/kumayacoin/explorer/templates"
	"github.com/magutae/kumayacoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
