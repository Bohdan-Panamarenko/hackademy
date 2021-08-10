unsigned int my_strlen(char *str)
{
    unsigned int len = 0;
    while (*(str++) != '\0')
    {
        len++;
    }
    return len;
}