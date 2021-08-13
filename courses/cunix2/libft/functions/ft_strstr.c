#include "../libft.h"
char *ft_strstr(const char *haystack, const char *needle)
{
    int needle_len = ft_strlen(needle);
    do
    {
        if (*haystack == *needle)
        {
            if (ft_strncmp(haystack, needle, needle_len) == 0)
            {
                return (char *)haystack;
            }
        }
    }
    while (*(haystack++) != '\0');
    return NULL;
}
