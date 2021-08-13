/* * ===================================================================================== *
 *       Filename:  libft.h
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  08/12/2021 01:41:14 PM
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  YOUR NAME (), 
 *   Organization:  
 *
 * =====================================================================================
 */
#ifndef _LIBFT_H_
#define _LIBFT_H_
    #include <stddef.h>
    #include <stdlib.h>
    const int UPPER_TO_LOWER_CASE_DIFFERENCE = 32;
    void ft_bzero(void *s, size_t n);

    char *ft_strdup(const char *s);

    unsigned int ft_strlen(const char *str);

    int ft_strncmp(const char *s1, const char *s2, size_t n);

    char *ft_strchr(const char *s, int c);

    char *ft_strrchr(const char *s, int c);

    int ft_isalpha (int c);

    int ft_isdigit(int c);
    
    int ft_isascii(int c);

    int ft_toupper(int c);

    int ft_tolower(int c);
#endif
