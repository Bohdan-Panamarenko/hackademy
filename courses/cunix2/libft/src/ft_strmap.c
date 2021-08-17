#include "../libft.h"

char *ft_strmap(char const *s, char (*f)(char))
{
    int s_len = ft_strlen((char *)s);
    char *new_s = (char *)malloc(sizeof(char) * (s_len + 1));
    new_s[s_len] = '\0';

    for (int i = 0; i < s_len; i++)
    {
        new_s[i] = f(s[i]); 
    }

    return new_s;
}
