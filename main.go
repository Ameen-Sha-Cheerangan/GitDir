// main.go is the entry point for gitdir.
// It delegates immediately to the cmd package — zero logic lives here.
package main

import "github.com/Ameen-Sha-Cheerangan/GitDir/cmd"

func main() {
	cmd.Execute()
}
