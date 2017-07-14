package main

	
import (
	"os"
	"fmt"
	"steamtrade"
)

func Usage() {
	fmt.Println("Usage: " + os.Args[0] + " /path/to/config")
}

func main() {
	if len(os.Args) != 2 {
		Usage();
		os.Exit(1);
	}
	
	config_path := os.Args[1]
	err := steamtrade.ReloadConfig(config_path)
	if err != nil {
		os.Exit(1)
	}
	
	steamtrade.StartServer()
	
}
