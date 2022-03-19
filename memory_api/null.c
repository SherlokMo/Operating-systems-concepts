#include <stdio.h>
#include <stdlib.h>


// will return a segmentation fault.
int main() {
    int *x = (int *) malloc(sizeof(int));
    x = NULL;
    printf("%d \n ", *x);
    free(x);
    return 0;
}