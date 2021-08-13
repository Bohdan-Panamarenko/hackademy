#include "../libft.h"
int ft_toupper(int c)
{
    if (c >= 'a' && c <= 'z')
    {
        return c - UPPER_TO_LOWER_CASE_DIFFERENCE; 
    }
    else
    {
        return c;
    }
}
