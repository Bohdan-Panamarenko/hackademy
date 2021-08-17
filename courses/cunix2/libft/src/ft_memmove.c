#include "../libft.h"

void *ft_memmove(void *dest, const void *src, size_t n)
{
    const char *src_str = (char *)src;
    char *dest_str = (char *)dest;
    char buffer[n];

    for (unsigned long i = 0; i < n; i++)
    {
        buffer[i] = src_str[i];
    }
    for (unsigned long i = 0; i < n; i++)
    {
        dest_str[i] = buffer[i];
    }

    return dest;
}
