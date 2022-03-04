package middlewares

import ("github.com/gominima/minima"
        "github.com/fatih/color")

func Logger(res *minima.Response,  req *minima.Request) {
	color.Blue("Method %s in route : %v", req.Method() , req.GetPathURL())
}