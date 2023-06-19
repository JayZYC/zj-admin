package util

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func PProf() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
