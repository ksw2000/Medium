// How to simply allocate 2D array at runtime (in heap segment)
#include <stdio.h>
#include <stdlib.h>

int** allocate2D(int rows, int cols) {
    // allocate memory for rows
    int** array2D = array2D = malloc(sizeof(int*) * rows);

    // allocate cols for each element in array2D
    int i;
    for (i = 0; i < rows; i++) {
        array2D[i] = malloc(sizeof(int) * cols);
    }
    return array2D;
}

void free2D(int** array2D, int rows) {
    int i;
    for (i = 0; i < rows; i++) {
        free(array2D[i]);
    }
    free(array2D);
}

int main() {
    int rows = 3;
    int cols = 5;
    int** array2D = allocate2D(rows, cols);
    // call malloc() 4 times

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

    free2D(array2D, rows);
    // call free() 4 times

    return 0;
}

