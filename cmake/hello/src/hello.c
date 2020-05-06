#include "hello.h"
#include "stdio.h"
#include "config.h"

#ifdef USE_MATH
#include "math.h"
#endif

int main(){
    printf("%s version is %d.%d\n", PROJECTNAME, HELLO_VERSION_MAJOR, HELLO_VERSION_MINOR);
    #ifdef USE_MATH
    printf("max value of %d and %d is %d\n", 3, 6, max(3,6));
    #endif
    return 0;
}