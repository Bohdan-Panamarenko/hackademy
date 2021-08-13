#include "../libft.h"
char *ft_strnstr(const char *big, const char *little, size_t len)
{
    int little_len = ft_strlen(little);
    for (size_t i = 0; i + little_len <= len; i++)
    {
        if (*big == *little)
        {
            if (ft_strncmp(big, little, little_len) == 0)
            {
                return (char *)big;
            }
        }
        if (*(big++) == '\0')
        {
            break;
        }
    }
    return NULL;
}
