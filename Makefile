all:
	gcc -shared -o ./c/libperson.so -fPIC ./c/person.c
	go build -o mainbin

