package main

import (
	"fmt"

	"github.com/vompressor/vplug/loader"
	"github.com/vompressor/vplug/vplugin"
)

func main() {
	var VPlugin *vplugin.VPlugin

	VPlugin, err := loader.Load("plugin_side/vpexample.so", "VPlugin")

	if err != nil {
		panic(err.Error())
	}

	println("Plugin name:")
	println(VPlugin.Info.Name)
	println()
	println("Plugin version")
	println(VPlugin.Info.Version)
	println()
	println("Plugin description")
	println(VPlugin.Info.Description)
	println()

	println("Function list:")
	for k, v := range VPlugin.FuncMap {
		fmt.Printf("%s:", k)
		for i, d := range v.InTypes {
			fmt.Printf(" %d:[%s],", i, d.String())
		}

		println()
	}

	println()

	VPlugin.FuncMap["hello"].Call()
	VPlugin.FuncMap["echo"].Call("hi")

	for k, v := range VPlugin.ValMap {
		fmt.Printf("%s:", k)
		println(v.Type.String())
		println(v.Val.(float64))
	}

}
