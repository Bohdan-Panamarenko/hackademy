#include "../libft.h"
void *ft_memchr(const void *s, int c, size_t n)
{
    const char *s_str = (char *)s;
    char ch = (char)c;
    for (size_t i = 0; i < n; i++)
    {
        if (s_str[i] == ch)
        {
            return (void *)(s + i);
        }
    }
    return NULL;
}
