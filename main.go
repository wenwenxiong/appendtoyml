package main

import (
	"fmt"
	"github.com/wenwenxiong/appendtoyml/cmd"
	"os"
)

func main() {
	rootCmd := cmd.NewAppendtoymlCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}