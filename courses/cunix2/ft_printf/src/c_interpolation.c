#include "../printf.h"

char *c_interpolation(char c, unsigned long width, int flags[4])
{
    flags[0] = 0;

    char *s = (char *)malloc(2 * sizeof(char));
    s[1] = '\0';
    s[0] = c;

    return s_float(s, width, flags);
}

