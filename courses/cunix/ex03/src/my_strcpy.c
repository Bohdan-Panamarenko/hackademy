char *my_strcpy(char *dest, const char *src)
{
    char *dest_to_return = dest; // save pointer to a start of dest
    while (*src != '\0') 
    {
        *(dest++) = *(src++);
    }
    *(dest) = '\0';
    return dest_to_return;
}