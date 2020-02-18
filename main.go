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

static ds GetData() {
	ds d;
	d.ptr = data;
	d.size = 1024;
	return d;
}
*/
import "C"
import (
	"log"
	"time"
	"unsafe"
)

func main() {
	t := time.Now()
	C.my_job()
	log.Println(time.Since(t))

	t = time.Now()
	i := 10000000
	for i > 0 {
		C.GetData()
		//log.Println(d)
		i = i-1
	}
	log.Println(time.Since(t))

}

//export my_callback
func my_callback(p unsafe.Pointer) {
	//log.Println((*C.ds)(p))
}
