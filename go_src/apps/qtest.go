package main

import (
	"fmt"
	"os"
)

func main() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("UserConfigDir error:", err)
		return
	}

	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("UserCacheDir error:", err)
		return
	}

	fmt.Println("UserConfigDir:", configDir)
	fmt.Println("UserCacheDir:", cacheDir)
}
