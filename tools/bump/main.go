package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const (
	configFileName = "CONFIG_FILE"
)

func main() {

	nBig, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		panic(err)
	}

	version := fmt.Sprintf("0.0.%d\n", nBig.Int64()) //#nosec G404
	d1 := []byte(version)

	filename := os.Getenv(configFileName)
	err = os.WriteFile(filename, d1, 0644) //#nosec G306

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", version)
}
