package test

import (
	"fmt"
	"sorm/v2/logo"
	"testing"
)

func TestLogo(t *testing.T) {
	logo.PrintLogo()
}

func TestColorPrint(t *testing.T) {
	fmt.Println(logo.Green("HELLO"))
}
