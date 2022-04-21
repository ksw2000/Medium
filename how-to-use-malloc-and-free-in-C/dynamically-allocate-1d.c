// How to allocate a 1d array at runtime (in heap segment)
#include <stdio.h>
#include <stdlib.h>

int* allocate1D(int n) {
    return malloc(sizeof(int) * n);
}

int main() {
    int n = 5;
    int* list = allocate1D(n);
    int i;
    for (i = 0; i < n; i++) {
        list[i] = i;
    }

    for (i = 0; i < n; i++) {
        printf("list[%d] = %d\n", i, list[i]);
    }

    // Output:
    // list[0] = 0
    // list[1] = 1
    // list[2] = 2
    // list[3] = 3
    // list[4] = 4
    
    free(list);
    return 0;
}