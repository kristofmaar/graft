package tensorflow

// #include <stdlib.h>
// #include "tensorflow/c/c_api_experimental.h"
import "C"
import (
	"errors"
	"unsafe"
)

func LoadPluggableDevice(filename string) error {
	status := C.TF_NewStatus()
	file := C.CString(filename)
	defer C.free(unsafe.Pointer(file))
	library := C.TF_LoadPluggableDeviceLibrary(file, status)
	defer C.free(unsafe.Pointer(library))
	if status != nil {
		defer C.TF_DeleteStatus(status)
	}
	if library == nil {
		return errors.New("warning: unable to load plugin")
	} else {
		return nil
	}
}
