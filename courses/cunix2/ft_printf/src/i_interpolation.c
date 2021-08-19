#include "../printf.h"

char *i_interpolation(int i, unsigned long width, int flags[4])
{
    char *s = ft_itoa(i);
    int empty_flags[4] = {0, 0, 0, 0 };
    unsigned long s_len = ft_strlen(s);

    if (*s == '-')
    {
        char *new_s = (char *)malloc(sizeof(char) * (s_len));
        
        for (unsigned int i = 1; i <= s_len; i++)
        {
            new_s[i - 1] = s[i];
        }

        free(s);
        s = new_s;
        s_len--;
    }

    if ((i < 0 || flags[2] || flags[3]) && !flags[0])
    {
        char *new_s = s_float(s, s_len + 1, empty_flags);
        
        if (i < 0)
        {
            new_s[0] = '-';
        }
        else if (flags[2])
        {
            new_s[0] = '+';
        }
        else 
        {
            new_s[0] = ' ';
        }

        s_len++;
        free(s);
        s = new_s;
    }

    if (width < s_len)
    {
        width = s_len;
    }

    char *new_s = s_float(s, width, flags);

    if (s != new_s)
    {
        free(s);
        s = new_s;
    }

    if ((i < 0 || flags[2] || flags[3]) && flags[0])
    {
        if (width == s_len)
        {
            width = s_len + 1;
            char *new_s = s_float(s, width, flags);
            free(s);
            s = new_s;
        }

 
        if (i < 0)
        {
            s[0] = '-';
        }
        else if (flags[2])
        {
            s[0] = '+';
        }
        else 
        {
            s[0] = ' ';
        }
    }
    
    return s;
}
