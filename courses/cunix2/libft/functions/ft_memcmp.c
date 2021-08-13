#include "../libft.h"
int ft_memcmp(const void *s1, const void *s2, size_t n)
{
    const char *s1_str = (const char *)s1;
    const char *s2_str = (const char *)s2;
    for (size_t i = 0; i < n; i++)
    {
        if (s1_str[i] < s2_str[i])
        {
            return 1;
        }
        else if (s1_str[i] > s2_str[i])
        {
            return -1;
        }
    }
    return 0;
}
