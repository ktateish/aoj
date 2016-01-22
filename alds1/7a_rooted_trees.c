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
	int parent;
	int child;
	int rsibling;
	int depth;
} node_t;


#define MAX_N 100010
int n;
node_t node[MAX_N];

const char *node_type_str(int id) {
	if (node[id].parent == -1) return "root";
	if (node[id].child == -1) return "leaf";
	else return "internal node";
}

void mark(int id, int d) {
	node[id].depth = d;
	if (node[id].child != -1) mark(node[id].child, d+1);
	if (node[id].rsibling != -1) mark(node[id].rsibling, d);
}

int main(int argc, char **argv)
{
	int i, j;
	memset(node, -1, sizeof(node_t) * MAX_N);
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		int id, k, left;
		scanf("%d %d", &id, &k);
		for (j = 0; j < k; j++) {
			if (j) {
				scanf("%d", &node[left].rsibling);
				node[node[left].rsibling].parent = id;
				left = node[left].rsibling;
			} else {
				scanf("%d", &node[id].child);
				node[node[id].child].parent = id;
				left = node[id].child;
			}
		}
	}

	for (i = 0; i < n; i++) {
		if (node[i].parent == -1) {
			mark(i, 0);
			break;
		}
	}

	for (i = 0; i < n; i++) {
		int child;
		const char *sp = "";

		printf("node %d: parent = %d, depth = %d, %s, [",
				i, node[i].parent, node[i].depth, node_type_str(i));
		child = node[i].child;
		while (child != -1) {
			printf("%s%d", sp, child);
			sp = ", ";
			child = node[child].rsibling;
		}
		printf("]\n");
	}
	return 0;
}

