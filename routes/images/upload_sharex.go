package images

import (
	"github.com/gominima/minima"
)

func UploadShareX(res *minima.Response, req *minima.Request) {
	res.Status(200)
	res.Send("Hello!")
}
