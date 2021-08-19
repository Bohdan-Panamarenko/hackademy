#include "../printf.h"

char *s_float(const char *s, unsigned long width, int flags[4])
{
    unsigned long s_len = ft_strlen(s);
    
    if (width <= s_len)
    {
        return (char *)s;
    }
    
    char *new_s = (char *)malloc(sizeof(char) * (width + 1));
    new_s[width] = '\0';
    char chr_for_fill = ' ';
    
    if (flags[0])
    {
        chr_for_fill = '0';
    }
    
    if (flags[1])
    {
        for (unsigned long i = s_len; i < width; i++)
        {
            new_s[i] = chr_for_fill;
        }
        
        for (unsigned long i = 0; i < s_len; i++)
        {
            new_s[i] = s[i];
        }
    }
    else
    {
        unsigned long start = width - s_len; // start of given string
        
        for (unsigned long i = 0; i < start; i++)
        {
            new_s[i] = chr_for_fill;
        }
        
        for (unsigned long i = start; i < width; i++)
        {
            new_s[i] = s[i - start];
        }
    }
    
    
    return new_s;
}
