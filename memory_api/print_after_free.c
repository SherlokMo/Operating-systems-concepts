#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[]) {

    int *data = (int *) malloc(100 * sizeof(int));

    data[100] = 100;

    free(data);

    printf("%d \n", data[100]);

}