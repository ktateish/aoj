#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

typedef struct node_st {
	int par;
	int left;
	int right;
} node_t;

enum {
	PREORDER,
	INORDER,
	POSTORDER,
};

#define MAX_N 64
int n;
node_t nd[MAX_N];

void print_if_match(int id, int o1, int o2) {
	if (o1 == o2) printf(" %d", id);
}

void dfs(int id, int order) {
	if (id == -1) return;
	print_if_match(id, order, PREORDER);
	dfs(nd[id].left, order);
	print_if_match(id, order, INORDER);
	dfs(nd[id].right, order);
	print_if_match(id, order, POSTORDER);
}

int main(int argc, char **argv)
{
	int i, id, l, r, root;

	memset(nd, -1, sizeof(node_t) * MAX_N);

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d %d %d", &id, &l, &r);
		nd[id].left = l;
		nd[id].right = r;
		if (l != -1) nd[l].par = id;
		if (r != -1) nd[r].par = id;
	}

	root = -1;
	for (i = 0; i < n; i++) {
		if (nd[i].par == -1) {
			root = i;
			break;
		}
	}
	puts("Preorder");
	dfs(root, PREORDER);
	puts("");
	puts("Inorder");
	dfs(root, INORDER);
	puts("");
	puts("Postorder");
	dfs(root, POSTORDER);
	puts("");
			
	return 0;
}

