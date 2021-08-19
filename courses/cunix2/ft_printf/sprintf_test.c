#include "./printf.h"
#include <stdio.h>

int main()
{
    char arr[256];
    int len = 0;

    len += ft_sprintf(arr, "(%s)", "hello");
    puts(arr);
    
    len += ft_sprintf(arr, "(%010d)", -23);
    puts(arr);

    len += ft_sprintf(arr, "(%+3i)", 1245);
    puts(arr);
}
