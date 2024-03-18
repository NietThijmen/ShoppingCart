package main

import (
	"github.com/NietThijmen/ShoppingCart/storage"
	"os"
	"os/exec"
)

func connect(item storage.SSHConfig) {
	// create a connection

	cmd := "ssh " + item.User + "@" + item.Host + " -p " + item.Port
	executed := exec.Command("sh", "-c", cmd)

	executed.Stdin = os.Stdin
	executed.Stdout = os.Stdout
	executed.Stderr = os.Stderr

	_ = executed.Run()
}
