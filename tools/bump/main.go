package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	configFileName = "CONFIG_FILE"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	version := fmt.Sprintf("0.0.%d\n", rand.Intn(1000)) //#nosec G404
	d1 := []byte(version)

	filename := os.Getenv(configFileName)
	err := os.WriteFile(filename, d1, 0644) //#nosec G306

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", version)
}
