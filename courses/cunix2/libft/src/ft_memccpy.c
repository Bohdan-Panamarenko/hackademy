#include "../libft.h"

void *ft_memccpy(void *dest, const void *src, int c, size_t n)
{
    const char *src_str = (char *)src;
    char *dest_str = (char *)dest;
    char ch = (char)c;

    for (unsigned long i = 0; i < n; i++)
    {
        dest_str[i] = src_str[i];
        if (dest_str[i] == ch)
        {
            return (dest + i + 1);
        }
    }

    return NULL;
}
