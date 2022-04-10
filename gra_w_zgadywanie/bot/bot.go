package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "/home/kuba/Repozytoria/sem4/golang-labs/gra_w_zgadywanie/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd.Stdout.Write([]byte("Hello")))
	fmt.Println(cmd.Stdout.Write([]byte("World")))
	cmd.Wait()

}
