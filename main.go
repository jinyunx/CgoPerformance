package main

/*
static char data[1024];
typedef struct ds {
	char *ptr;
	int size;
} ds;
extern void my_callback(void*);
static void my_job() {
  for (int i = 0; i < 10000000; ++i) {
	ds d;
	d.ptr = data;
	d.size = 1024;
	my_callback(&d);
  }
}
static void GetData() {
	ds d;
	d.ptr = data;
	d.size = 1024;
	return;
}
*/
import "C"
import (
	"log"
	"syscall"
	"time"
	"unsafe"
)

var (
	dll = syscall.MustLoadDLL("dll.dll")

	procDllGetData = dll.MustFindProc("DllGetData")
	procDllJob     = dll.MustFindProc("DllJob")
	procDllSet     = dll.MustFindProc("DllSet")
)

func main() {
	//// cgo start
	t := time.Now()
	C.my_job()
	log.Println("cgo callback:", time.Since(t))

	t = time.Now()
	i := 10000000
	for i > 0 {
		C.GetData()
		//log.Println(d)
		i = i - 1
	}
	log.Println("cgo call:", time.Since(t))

	//// syscall
	callback := syscall.NewCallback(go_callback)
	_, _, _ = procDllSet.Call(callback)
	t = time.Now()
	_, _, _ = procDllJob.Call()
	log.Println("syscall callback:", time.Since(t))

	t = time.Now()
	i = 10000000
	for i > 0 {
		_, _, _ = procDllGetData.Call()
		//log.Println(d)
		i = i - 1
	}
	log.Println("syscall call:", time.Since(t))

}

func go_callback(p unsafe.Pointer) (ret uintptr) {
	//log.Println(p)
	ret = 0
	return ret
}

//export my_callback
func my_callback(p unsafe.Pointer) {
	//log.Println((*C.ds)(p))
}
