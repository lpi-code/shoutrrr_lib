package main

import (
	"C"

	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/containrrr/shoutrrr/pkg/types"
)

// Send notifications using a supplied url and message
//
import "unsafe"

//export Send
func Send(rawURL *C.char, message *C.char) *C.char {
	var defaultRouter = router.ServiceRouter{}
	service, err := defaultRouter.Locate(C.GoString((*C.char)(unsafe.Pointer(rawURL))))
	if err == nil {
		err = service.Send(C.GoString((*C.char)(unsafe.Pointer(message))), &types.Params{})
	}
	// log the message to the console
	if err != nil {
		return C.CString(err.Error())
	}
	return C.CString("")
}

func main() {}
