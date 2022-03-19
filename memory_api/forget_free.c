#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[]) {
    int *x = (int *) malloc(sizeof(int));
    *x = 2022;
    printf("%d \n", *x); // printing value of x
    printf("%p \n", x);  // printing memory adress of x (pointer location)

    return 0;
}