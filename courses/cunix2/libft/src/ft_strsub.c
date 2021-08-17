#include "../libft.h"

char *ft_strsub(char const *s, unsigned int start, size_t len)
{
    unsigned int s_len = ft_strlen(s);
    
    if (start >= s_len) 
    {
        len = 0;
    }
    else if (s_len - start < len)
    {
        len = s_len - start;
    }

    char *new_s = (char *)malloc(sizeof(char) * (len + 1));
    new_s[len] = '\0';

    for (size_t i = 0; i < len; i++)
    {
        new_s[i] = s[i + start];
    }

    return new_s;
    
}
