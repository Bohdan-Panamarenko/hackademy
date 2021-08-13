#include "../libft.h"
void *ft_memcpy(void *dest, const void *src, size_t n)
{
    const char *src_str = (char *)src;
    char *dest_str = (char *)dest;
    for (size_t i = 0; i < n; i++)
    {
        dest_str[i] = src_str[i];
    }
    return dest;
}
