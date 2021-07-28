package main

import (
	"github.com/vompressor/vplug/vplugin"
)

func main() {}

// REQUIREMENT!! Define plugin info
var VPlugin = vplugin.NewVPlugin("vpexample", "v0.0.0", "example project").
	AddVPVal("pi", 3.14).
	AddVPFunc("hello", func() error { println("hello"); return nil }).
	AddVPFunc("echo", func(str string) error { println(str); return nil })
