#include <stdio.h>
#include <assert.h>
#include <pthread.h>
#include "common.h"
#include "common_threads.h"


static volatile int counter = 0;


void *mythread(void *arg) {
    printf("%s begin \n", (char *) arg);

    for(int i = 0; i < 1e7; i++) {
        counter++;
    }
    printf("%s done \n", (char *) arg);
    return NULL;
}


int main(int argc, char *argv[]) {
    pthread_t t1, t2;

    printf("main: beign (counter = %d) \n", counter);
    Pthread_create(&t1, NULL, mythread, "A");
    Pthread_create(&t2, NULL, mythread, "B");

    Pthread_join(t1, NULL);
    Pthread_join(t2, NULL);
    printf("main: done with both (counter = %d) \n", counter);

    return 0;
}