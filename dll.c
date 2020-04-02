// gcc -c dll.c
// gcc -shared dll.o -o dll.dll

static char data[1024];

typedef struct ds {
	char *ptr;
	int size;
} ds;

typedef void (__stdcall *dll_callback)(void*);

dll_callback _cb;
void DllSet(dll_callback cb) {
    _cb = cb;
}

void DllJob() {
  for (int i = 0; i < 10000000; ++i) {
	ds d;
	d.ptr = data;
	d.size = 1024;
	_cb(&d);
  }
}

void DllGetData() {
	ds d;
	d.ptr = data;
	d.size = 1024;
	return;
}
