#include "../libft.h"

unsigned int ft_strlen(const char *str)
{
    unsigned int len = 0;

    while (*(str++) != '\0')
    {
        len++;
    }

    return len;
}
