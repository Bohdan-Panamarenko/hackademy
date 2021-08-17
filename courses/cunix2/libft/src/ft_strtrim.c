#include "../libft.h"

char *ft_strtrim(char const *s)
{
    unsigned int s_len = ft_strlen(s);
    unsigned int i = 0;

    while ((s[i] == ' ' || s[i] == '\t' || s[i] == '\n') && s[i] != '\0') 
    {
        i++;
    }
    s = &s[i];
    s_len -= i;

    if (s_len != 0)
    {
        while (s[s_len - 1] == ' ' || s[s_len - 1] == '\t' || s[s_len - 1] == '\n')
        {
            s_len--;
        }
    }

    char *new_s = (char *)malloc(sizeof(char) * (s_len + 1));
    new_s[s_len] = '\0';

    ft_memcpy(new_s, s, s_len);

    return new_s;
}
