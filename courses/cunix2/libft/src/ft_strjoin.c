#include "../libft.h"

char *ft_strjoin(char const *s1, char const *s2)
{
    size_t s1_len = ft_strlen(s1);
    size_t new_s_len = s1_len + ft_strlen(s2); 
    char *new_s = (char *)malloc(sizeof(char) * (new_s_len + 1));
    new_s[new_s_len] = '\0';

    for (size_t i = 0; i < s1_len; i++)
    {
        new_s[i] = s1[i];
    }

    for (size_t i = s1_len; i < new_s_len; i++)
    {
        new_s[i] = s2[i - s1_len];
    }

    return new_s;
}
