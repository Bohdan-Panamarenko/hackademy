#include "../libft.h"

char *ft_strchr(const char *s, int c)
{
    char chr = (char)c;

    do 
    {
        if (*s == chr)
        {
            return (char *)s;
        }
    }
    while (*(s++) != '\0');

    return 0;
}
