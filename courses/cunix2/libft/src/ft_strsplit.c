#include "../libft.h"

char **ft_strsplit(char const *s, char c)
{
    unsigned long s_len = ft_strlen(s);
    unsigned long i = 0;
    int num_of_str = 0;
    long start_of_str = 0;

    while (i < s_len)
    {
        if (start_of_str && s[i] == c)
        {
            start_of_str = 0;
            num_of_str++;
        }
        else if (s[i] != c)
        {
            start_of_str = 1;
        }

        i++;
    }

    if (start_of_str)
    {
        num_of_str++;
    }

    char **strs = (char **)malloc(sizeof(char *) * (num_of_str + 1));
    strs[num_of_str] = NULL;

    i = 0;
    unsigned long str_len = 0;
    num_of_str = 0;
    start_of_str = -1;

    while (i < s_len)
    {
        if (start_of_str == -1)
        {
            if (s[i] != c)
            {
                start_of_str = i;
                str_len++;
            }
        }
        else
        {
            if (s[i] == c)
            {
                strs[num_of_str] = (char *)malloc(sizeof(char) * (str_len + 1));
                strs[num_of_str][str_len] = 0;
                ft_memcpy(strs[num_of_str], &s[start_of_str], str_len);
                
                start_of_str = -1;
                str_len = 0;
                num_of_str++;
            }
            else
            {
                str_len++;
            }
        }

        i++;
    }

    if (start_of_str != -1)
    {
        strs[num_of_str] = (char *)malloc(sizeof(char) * (str_len + 1));
        strs[num_of_str][str_len] = 0;
        ft_memcpy(strs[num_of_str], &s[start_of_str], str_len);
    }

    return strs;
}
