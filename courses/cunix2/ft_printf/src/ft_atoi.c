#include "../printf.h"

int ft_atoi(const char *nptr)
{

    const int int_to_char_offset = 48; // offset from digit to asci digit char
    int num = 0, minus = 1;

    if (*nptr == '-') // if there is minus in nptr, make num negative
    {
        minus = -1;
        nptr++;
    }

    for (int i = 0; nptr[i] <= '9' && nptr[i] >= '0'; i++) // reading digits from nptr and producing num
    {
        num *= 10;
        num += (nptr[i] - int_to_char_offset) * minus;
    }

    return num;
}
