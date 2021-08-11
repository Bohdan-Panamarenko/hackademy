//#include <string.h>
int my_strcmp(char *str1, char *str2) 
{
    while (*str1 != '\0' && *str2 != '\0') // lexical comparison
    {
        if (*str1 > *str2)
        {
            return 1;
        }
        else if (*str1 < *str2)
        {
            return -1;
        }
        str1++;
        str2++;
    }
    if (*str1 != '\0' && *str2 == '\0') // length comparison
    {
        return 1;
    }
    else if (*str1 == '\0' && *str2 != '\0')
    {
        return -1;
    }
    else
    {
        return 0;
    }
}