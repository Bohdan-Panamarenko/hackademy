#include "../libft.h"

void ft_bzero(void *s, size_t  n)
{
    char *str = (char *)s;
    for (size_t i = 0; i < n; i++)
    {
        str[i] = '\0';
    }
}
