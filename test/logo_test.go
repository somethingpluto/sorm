package test

import (
	"fmt"
	logo2 "sorm/logo"
	"testing"
)

func TestLogo(t *testing.T) {
	logo2.PrintLogo()
}

func TestColorPrint(t *testing.T) {
	fmt.Println(logo2.Green("HELLO"))
}
