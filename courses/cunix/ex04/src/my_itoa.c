#include <stdlib.h>
char *my_itoa(int nmb)
{
    const int int_to_char_offset = 48; // offset from digit to asci digit char
    int str_len = 1, start = 0;
    if (nmb < 0) // if number is negativ then make space for '-' and start digits from index 1
    {
        nmb = -nmb;
        start = 1;
        str_len++;
    }
    for (int x = nmb; x / 10 != 0; x /= 10) // determining the length of string
    {
        str_len++;
    }
    char *str = (char *)malloc((str_len + 1) * sizeof(char)); // creating the string
    str[str_len] = '\0';
    if (start) // placing the '-'
    {
        str[0] = '-';
    }
    for (int i = str_len - 1; i >= start; i--) // placing the digits
    {
        str[i] = nmb % 10 + int_to_char_offset;
        nmb /= 10;
    }
    return str;
}
