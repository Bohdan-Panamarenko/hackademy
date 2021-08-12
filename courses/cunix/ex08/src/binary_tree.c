#include "../include/binary_tree.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
node_t  *allocnode()
{
    node_t *node = (node_t *)malloc(sizeof(node_t));
    node->right = NULL;
    node->left = NULL;
    return node;
}
node_t *insert(node_t *root, char *key, void *data)
{
    if (root)
    {
        if (strcmp(key, root->key) < 0)
        {
            root->left = insert(root->left, key, data);
        }
        else
        {
            root->right = insert(root->right, key, data);
        }
        return root;
    }
    else
    {
        node_t *node = allocnode();
        node->key = key;
        node->data = data;
        return node;
    }
}
void print_node(node_t *node)
{
    if (node)
    {
        printf("%s | %s\n", node->key, (char *)node->data);
    }
}
void visit_tree(node_t *node, void (*fp)(node_t *root))
{
    if (!node)
    {
        return;
    }
    visit_tree(node->left, fp);
    visit_tree(node->right, fp);
    fp(node);
}
void destroy_tree(node_t *node, void (*fdestroy)(node_t *root))
{
    visit_tree(node, fdestroy);
}
