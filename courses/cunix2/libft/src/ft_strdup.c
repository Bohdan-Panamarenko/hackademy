#include "../libft.h"

char *ft_strdup(const char *s)
{
    int len = ft_strlen((const char *)s);
    char *str = (char *)malloc(sizeof(char) * len + 1);

    for (int i = 0; i <= len; i++)
    {
        str[i] = s[i];
    }

    return str;
}

