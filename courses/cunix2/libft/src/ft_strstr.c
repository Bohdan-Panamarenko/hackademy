#include "../libft.h"

char *ft_strstr(const char *haystack, const char *needle)
{
    unsigned long needle_len = ft_strlen(needle);
    unsigned long haystack_len = ft_strlen(haystack);

    for (int i = 0; i + needle_len <= haystack_len; i++) 
    {
        if (ft_strncmp(&haystack[i], needle, needle_len) == 0)
        {
            return (char *)&haystack[i];
        }
    }

    return NULL;
}
