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
	int left, right;
	int sib;
	int deg;
	int dep;
	int hei;
} node_t;

#define MAX_N 30

int n;
node_t nd[MAX_N];

const char *node_type_str(int id) {
	if (nd[id].par == -1) return "root";
	if (nd[id].left == -1 && nd[id].right == -1) return "leaf";
	else return "internal node";
}

void dfs(int id, int d) {
	int l, r;

	if (id == -1) return;
	l = nd[id].left;
	r = nd[id].right;

	nd[id].dep = d;
	dfs(l, d+1);
	dfs(r, d+1);

	nd[id].hei = 0;
	if (l != -1) {
		nd[id].hei = nd[l].hei + 1;
	}
	if (r != -1) {
		if (nd[id].hei <= nd[r].hei) {
			nd[id].hei = nd[r].hei + 1;
		}
	}
}

void pnode(int id) {
	printf("node %d: parent = %d, sibling = %d, degree = %d, "
	       "depth = %d, height = %d, %s\n",
	       id, nd[id].par, nd[id].sib, nd[id].deg, nd[id].dep,
	       nd[id].hei, node_type_str(id));
}

int main(int argc, char **argv)
{
	int i, id, l, r;

	memset(nd, -1, sizeof(node_t) * MAX_N);
	scanf("%d", &n);

	for (i = 0; i < n; i++) {
		scanf("%d %d %d", &id, &l, &r);
		nd[id].left = l;
		nd[id].right = r;
		nd[id].deg = 0;
		if (l != -1 && r != -1) {
			nd[l].sib = r;
			nd[r].sib = l;
		}
		if (l != -1) {
			nd[l].par = id;
			nd[id].deg++;
		}
		if (r != -1) {
			nd[r].par = id;
			nd[id].deg++;
		}
	}

	for (i = 0; i < n; i++) {
		if (nd[i].par == -1) {
			dfs(i, 0);
			break;
		}
	}

	for (i = 0; i < n; i++) {
		pnode(i);
	}

	return 0;
}

