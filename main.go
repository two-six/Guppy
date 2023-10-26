package main

import (
	"time"

	pkgterm "github.com/pkg/term"

	"projects/twpsx/guppy/engine"
)

func main() {
	t, err := pkgterm.Open("/dev/tty", pkgterm.RawMode)
	if err != nil {
		panic(err)
	}
	defer t.Restore()
	engine, err := engine.New()
	if err != nil {
		panic(err)
	}
	if err = engine.RunCommand("vSplit"); err != nil {
		panic(err)
	}
	if err = engine.RunCommand("hSplit"); err != nil {
		panic(err)
	}
	engine.Draw()
	time.Sleep(time.Second * 3)
}
