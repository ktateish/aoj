#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_N 1000010

typedef struct node_st {
	int v;
	struct node_st *next;
} node_t;

node_t *G[MAX_N];
int n;
int connected[MAX_N];
int color;

node_t *add_node(node_t *node, int v) {
	node_t *new = malloc(sizeof(node_t));
	new->v = v;
	new->next = node;
	return new;
}

void add(int u, int v) {
	G[u] = add_node(G[u], v);
	G[v] = add_node(G[v], u);
}

void dfs(int v) {
	node_t *p;
	if (connected[v]) return;
	connected[v] = color;
	for (p = G[v]; p != NULL; p = p->next) {
		dfs(p->v);
	}
}

int main(int argc, char **argv)
{
	int i, m, q, u, v;

	scanf("%d %d", &n, &m);
	for (i = 0; i < m; i++) {
		scanf("%d %d", &u, &v);
		add(u, v);
	}

	for (i = 0; i < n; i++) {
		color = i + 1;
		dfs(i);
	}

	scanf("%d", &q);
	for (i = 0; i < q; i++) {
		scanf("%d %d", &u, &v);
		if (connected[u] == connected[v]) puts("yes");
		else puts("no");
	}

	return 0;
}
