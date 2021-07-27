package main

import (
	"encoding/json"

	"github.com/vompressor/vplug/plugin"
)

func main() {}

// REQUIREMENT!! Define plugin info
var PluginInfo = plugin.PluginInfo{
	Name:    "vpexample",
	Version: "v0.0.0",
}

// Add funcs
var Funcs = plugin.NewNTFuncs(
	[]plugin.NTFuncHelper{

		// Print Hello
		plugin.NewNTFunc("Hello", func() error { println("hello"); return nil }),

		// Encode struct to json
		plugin.NewNTFunc("Encode", func(data interface{}, callback func([]byte)) error {
			d, err := json.Marshal(data)
			if err != nil {
				return err
			}
			callback(d)
			return nil
		}),
	},
)
