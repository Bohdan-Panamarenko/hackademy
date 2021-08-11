#include <stdlib.h>
#include <stdio.h>
#include "../include/linked_list.h"

node_t  *list_create(void *data)
{
    node_t *node = (node_t *)malloc(sizeof(node_t));
    node->data = data;
    node->next = NULL;
    return node;
}

void    *list_pop(node_t **head)
{
    node_t *prev = *head;
    if (prev->next->next != NULL)
    {
        while (prev->next->next != NULL)
        {
            prev = prev->next;
        }
    }
    void *data = prev->next->data;
    free(prev->next);
    prev->next = NULL;
    return data;
}

void    list_destroy(node_t **head, void (*fp)(void *data))
{
    if (!*head)
    {
        return;
    }
    else {
        node_t *list = *head;
        while (list->next != NULL)
        {
            void *data = list_pop(head);
            fp(data);
        }
        fp((*head)->data);
        free(*head);
        head = NULL;
    }
}

void    list_push(node_t *head, void *data)
{
    node_t *last = head;
    while (last->next != NULL)
    {
        last = last->next;
    }
    last->next = list_create(data);
}

void    list_unshift(node_t **head, void *data)
{
    if (*head)
    {
        *head = list_create(data);
    }
    else
    {
        node_t *old_list = *head;
        *head = list_create(data);
        (*head)->next = old_list;
    }
}


void    *list_shift(node_t **head)
{
    if (!*head)
    {
        return NULL;
    }
    else {
        node_t *new_list = NULL;
        if ((*head)->next)
        {
            new_list = (*head)->next;
            (*head)->next = NULL;
        }
        
        void *data = (*head)->data;
        free(*head);
        *head = new_list;
        return data;
    }
}
void    *list_remove(node_t **head, int pos)
{
    if (pos < 1 || !*head)
    {
        return NULL;
    }
    node_t *prev = *head, *curr_pos = (*head)->next;
    if (pos == 1)
    {
        return list_shift(head);
    }
    else
    {
        int i = 0;
        void *data = NULL;
        while (i != pos && curr_pos->next != NULL)
        {
            prev = curr_pos;
            curr_pos = curr_pos->next;
            i++;
        }
        if (curr_pos->next == NULL)
        {
            if (i == pos)
            {
                data = curr_pos->data;
                free(prev->next);
                prev->next = NULL;
            }
            else
            {
                return NULL;
            }
        }
        else
        {
            node_t *next = curr_pos->next;
            data = curr_pos->data;
            free(prev->next);
            prev->next = next;
        }
            return data;
    }
}

void    list_print(node_t *head)
{
    if (!head) {
        return;
    }
    else {
        do 
        {
            printf("%s\n", (char*)head->data);
            head = head->next;
        }
        while (head);
    }
}
void    list_visitor(node_t *head, void (*fp)(void *data))
{
    if (!head) {
        return;
    }
    else {
        do 
        {
            fp(head->data);
            head = head->next;
        }
        while (head);
    }
}