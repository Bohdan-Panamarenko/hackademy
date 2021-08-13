#include "../libft.h"
int ft_strncmp(const char *s1, const char *s2, size_t n)
{
    size_t i;
    for (i = 0; i != n && s1[i] != '\0' && s2[i] != '\0' ; i++) // lexical comparison
    {
        if (s1[i] > s2[i])
        {
            return 1;
        }
        else if (s1[i] < s2[i])
        {
            return -1;
        }
    }
    if (i == n)
    {
        return 0;
    }
    if (s1[i] != '\0' && s2[i] == '\0') // length comparison
    {
        return 1;
    }
    else if (s1[i] == '\0' && s2[i] != '\0')
    {
        return -1;
    }
    else
    {
        return 0;
    }
}
