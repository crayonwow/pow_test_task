package equix

/*
#cgo CFLAGS: -I${SRCDIR}/equix/include
#cgo LDFLAGS: -L${SRCDIR}/equix -lequix
#include <equix.h>
*/
import "C"
import "unsafe"

// import _cgopackage "runtime/cgo"

func Solve(challenge []byte) []byte {
	cChallenge := unsafe.Pointer(&challenge[0])
	cSolution := &C.equix_solution{}
	// cCtx := &C.equix_ctx{}
	C.equix_solve(nil, cChallenge, C.size_t(len(challenge)), cSolution)
	// defer C.equix_free(cCtx)
	return C.GoBytes(unsafe.Pointer(&cSolution.idx[0]), C.int(C.EQUIX_MAX_SOLS))
}

