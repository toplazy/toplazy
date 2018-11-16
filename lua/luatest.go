package main

import (
	"github.com/yuin/gopher-lua"
)




func luademo ()  {

	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("lua/test.lua"); err != nil {
		panic(err)
	}

}