package main

import (
	"fmt"
	"os"

	"github.com/sacas-network/sacas/app"
	"github.com/cosmos/cosmos-sdk/server"
)

func main() {
	// Buat instance aplikasi
	myApp := app.NewApp("SacasApp", nil)

	// Jalankan server
	if err := server.Start(myApp); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}
}

