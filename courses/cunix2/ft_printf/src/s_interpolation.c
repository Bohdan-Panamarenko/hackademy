#include "../printf.h"

char *s_interpolation(const char *s, unsigned long width, int flags[4])
{  
    if (!s)
    {
        s = "(null)";
    }

    flags[0] = 0;

    return s_float(s, width, flags);
}
