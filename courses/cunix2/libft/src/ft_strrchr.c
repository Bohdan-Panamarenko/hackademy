#include "../libft.h"

char *ft_strrchr(const char *s, int c)
{
    const char *s_end = s;
    char chr = (char)c;
    
    while (*s_end != '\0')
    {
        s_end++;
    }

    do 
    {
        if (*s_end == chr)
        {
            return (char *)s_end;
        }
    }
    while (s_end-- != s);
    
    return 0;
}
