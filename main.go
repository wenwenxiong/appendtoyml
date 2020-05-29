package main

import (
	"fmt"
	"os"

	"github.com/wenwenxiong/appendtoyml/cmd"
)

func main() {
	rootCmd := cmd.NewAppendtoymlCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}