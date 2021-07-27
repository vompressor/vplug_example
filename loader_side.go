package main

import (
	"fmt"

	"github.com/vompressor/vplug/loader"
	"github.com/vompressor/vplug/plugin"
)

func main() {
	var funcs *plugin.NTFuncs
	var info *plugin.PluginInfo

	// Load plugin
	funcs, info, err := loader.Load("plugin_side/vpexample.so", "Funcs")

	if err != nil {
		panic(err)
	}

	// Read plugin info
	println("plugin name:")
	println(info.Name)
	println("plugin info:")
	println(info.Version)
	println()
	println("funcs:")

	// Print funcs name and input types from plugin
	for k, v := range *funcs {
		println(k)
		println("in type:")
		for _, t := range v.InTypes {
			print(t.String() + ", ")
		}
		println()
	}
	println()

	// Get func Hello
	Hello := (*funcs)["Hello"]

	// Call Hello
	Hello.Call()

	// Get func Encode
	Encode := (*funcs)["Encode"]

	// Call func Encode
	Encode.Call(info, func(ret []byte) {
		fmt.Printf("%s\n", ret)
	})
}
