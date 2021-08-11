#include <stdio.h>
int my_puts(const char *s)
{
    while (*s != '\0') // reading string till the end and placing each character
    {
        putchar(*s);
        s++;
    }
    putchar('\n'); // placing end of line
    return 0;
}