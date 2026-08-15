// main.go is the entry point for gitdir.
// It delegates immediately to the cmd package — zero logic lives here.
package main

import "github.com/ameen/gitdir/cmd"

func main() {
	cmd.Execute()
}
