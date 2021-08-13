#include "../libft.h"
void *ft_memset(void *s, int c, size_t n)
{
    char *str = (char *)s;
    char chr = (char)c;
    for (size_t i = 0; i < n; i++)
    {
        str[i] = chr;
    }
    return s;
}