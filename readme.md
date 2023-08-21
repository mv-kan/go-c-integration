# Go C integration 

This is minimal example of integration of existing C code base into Go 

In C we have function called `person_is_adult()` that accepts `struct person` and returns int value 1 or 0 and this function is integrated with Go

Here the quick rundown of what is done 

```
// #cgo CFLAGS: -I./c
// #cgo LDFLAGS: -L./c -lperson -Wl,-rpath=./c
// #include "./c/person.h"
// #include <stdlib.h>
import "C"
```

`#cgo CFLAGS: -I./c` - sets CFLAGS and sets directory to `./c` 
`#cgo LDFLAGS: -L./c -lperson -Wl,-rpath=./c` - in makefile we compile libperson.o object and put it into `./c` directory, this line imports this object from `./c` directory
`#include "./c/person.h"` - includes header files with all definitions
`#include <stdlib.h>` - includes `free()` function (and others but we dont use them)