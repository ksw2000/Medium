// How to allocate contiguous 2D array at runtime (in heap segment)
#include <stdio.h>
#include <stdlib.h>

int** allocate2D(int rows, int cols) {
    size_t overhead = sizeof(int*) * rows;
    void** array = malloc(overhead + sizeof(int) * rows * cols);
    int i;
    for (i = 0; i < rows; i++) {
        array[i] = array + overhead + i * cols;
    }
    return (int**)array;
}

void free2D(int** array) {
    free(array);
}

int main() {
    int rows = 3;
    int cols = 5;
    int** array2D = allocate2D(rows, cols);
    // call malloc() only 1 time

    // TEST
    int i, j, count = 0;
    for (i = 0; i < rows; i++) {
        for (j = 0; j < cols; j++) {
            array2D[i][j] = count;
            count++;
        }
    }

    for (i = 0; i < rows; i++) {
        for (j = 0; j < cols; j++) {
            printf("%2d ", array2D[i][j]);
        }
        printf("\n");
    }
    
    // Output:
    //  0  1  2  3  4
    //  5  6  7  8  9
    // 10 11 12 13 14

    free2D(array2D);
    // call free() only 1 time

    return 0;
}
