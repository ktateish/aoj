#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define HSIZE 10009

typedef enum {
	NEXT_A,
	NEXT_C,
	NEXT_G,
	NEXT_T,
	NEXT_NUL, // NUL
} next_e;

int getidx(char c) {
	switch (c) {
	case 'A': return NEXT_A; break;
	case 'C': return NEXT_C; break;
	case 'G': return NEXT_G; break;
	case 'T': return NEXT_T; break;
	default:
		  fprintf(stderr, "Invalid char %c\n", c);
		  exit(1);
	}
}

typedef struct node_st {
	struct node_st *next[sizeof(next_e)];
} node_t;

void *xmalloc(size_t s) {
	void *p = malloc(s);
	if (p == NULL) {
		fprintf(stderr, "malloc failed\n");
		exit(1);
	}
	return p;
}

node_t *node_new() {
	node_t *p = xmalloc(sizeof(node_t));
	memset(p, 0, sizeof(node_t));
	return p;
}

node_t root_node, term_node;
node_t *dict = &root_node;
node_t *nul = &term_node;

void node_insert(node_t *n, const char *s) {
	char c = s[0];
	int i;
	if (c == '\0') {
		n->next[NEXT_NUL] = nul;
		return;
	}
	i = getidx(c);
	if (!n->next[i]) n->next[i] = node_new();
	node_insert(n->next[i], s + 1);
}

int node_search(node_t *n, const char *s) {
	char c = s[0];
	int i;
	if (c == '\0') {
		return n->next[NEXT_NUL] == nul;
	}
	i = getidx(c);
	if (n->next[i]) return node_search(n->next[i], s + 1);
	else return 0;
}

void insert(const char *s) {
	node_insert(dict, s);
}

int search(const char *s) {
	return node_search(dict, s);
}

int main(int argc, char **argv)
{
	int i, n;
	char cmd[8], s[16];

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%s %s", cmd, s);
		if (strcmp(cmd, "insert") == 0) {
			insert(s);
		} else {
			if (search(s)) puts("yes");
			else puts("no");
		}
	}
		
	return 0;
}

