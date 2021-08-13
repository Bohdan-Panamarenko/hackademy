#include "../libft.h"
int ft_tolower(int c)
{
    if (c >= 'A' && c <= 'Z')
    {
        return c + UPPER_TO_LOWER_CASE_DIFFERENCE; 
    }
    else
    {
        return c;
    }
}
