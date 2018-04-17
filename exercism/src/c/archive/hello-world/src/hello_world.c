#include <stdio.h>
#include <stddef.h>
#include <string.h>

#include "hello_world.h"

void hello(char *buffer, const char *name) {
    if (name == NULL) {
        sprintf(buffer, "%s", "Hello, World!");
        return;
    }

    sprintf(buffer, "%s, %s!", "Hello", name);
}
