package mymodule

import _ "github.com/valyala/fasthttp"

//ShowDep returns info about mymodule version and it's dependencies
func ShowDep() string {
	return "mymodule v0.0.1, fasthttp v1.15.1"
}
