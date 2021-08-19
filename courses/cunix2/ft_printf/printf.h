#include <stdlib.h>
#include <stdarg.h>
#include <unistd.h>

unsigned long ft_strlen(const char *s);

void ft_printf(const char *format, ...);

int ft_atoi(const char *nptr);

char *ft_itoa(int nmb);

char *s_float(const char *s, unsigned long width, int flags[4]);

char *c_interpolation(char c, unsigned long width, int flags[4]);

char *s_interpolation(const char *s, unsigned long width, int flags[4]);

char *i_interpolation(int i, unsigned long width, int flags[4]);

int ft_sprintf(char *arr, const char *format, ...);

