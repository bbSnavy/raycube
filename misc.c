#include <stdlib.h>
#include <stdio.h>

void    iterate_array_f32(float *arr, int len)
{
    int x;

    for (x = 0; x < len; x++)
    {
        float   *ptr = &arr[x];

        printf("\n");
        printf("%d -> [%f, %f, %f]\n", x, ptr[0], ptr[1], ptr[2]);
        ptr[0] *= 2.0;
        ptr[1] *= 2.0;
        printf("%d -> [%f, %f, %f]\n", x, ptr[0], ptr[1], ptr[2]);
    }
}
