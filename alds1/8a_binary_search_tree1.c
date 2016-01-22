#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

typedef struct node_st {
	int key;
	struct node_st *left, *right;
} node_t;

node_t *root;

node_t *node_new(int x) {
	node_t *new = malloc(sizeof(node_t));
	if (new == NULL) exit(1);
	new->key = x;
	new->left = new->right = NULL;
	return new;
}

node_t *insert(node_t *nd, int x) {
	if (!nd) return node_new(x);
	if (x < nd->key) nd->left = insert(nd->left, x);
	else if (nd->key < x) nd->right = insert(nd->right, x);
	return nd;
}

enum {
	PREORDER,
	INORDER,
};

void print_tree(node_t *nd, int order) {
	if (!nd) return;
	if (order == PREORDER) printf(" %d", nd->key);
	print_tree(nd->left, order);
	if (order == INORDER) printf(" %d", nd->key);
	print_tree(nd->right, order);
}

int main(int argc, char **argv)
{
	int i, n, x;
	char cmd[128];

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%s", cmd);
		if (!strcmp("insert", cmd)) {
			scanf("%d", &x);
			root = insert(root, x);
		} else if (!strcmp("print", cmd)) {
			print_tree(root, INORDER);
			puts("");
			print_tree(root, PREORDER);
			puts("");
		}
	}

	return 0;
}

