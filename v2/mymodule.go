package mymodule

import _ "github.com/valyala/fasthttp"

//ShowDep returns info about mymodule version and it's dependencies
func ShowDep() string {
	return "mymodule v2.0.2, fasthttp v1.21"
}
