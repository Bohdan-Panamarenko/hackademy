#include "../libft.h"
div_t ft_div(int numer, int denom)
{
    div_t div;
    div.quot = numer / denom;
    div.rem = numer % denom;
    return div;
}
