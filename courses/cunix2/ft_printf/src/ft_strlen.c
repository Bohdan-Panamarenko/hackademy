#include "../printf.h"

unsigned long ft_strlen(const char *s)
{
    unsigned long len = 0;

    while (*s != '\0')
    {
        len++;
        s++;
    }

    return len;
}
