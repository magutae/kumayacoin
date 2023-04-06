package main

import (
	"github.com/magutae/kumayacoin/cli"
	"github.com/magutae/kumayacoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
