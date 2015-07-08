#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;

typedef struct node_st {
	int val;
	struct node_st *prev;
	struct node_st *next;
} node_t;

typedef struct {
	node_t *head;
	node_t *tail;
	node_t *fl;
} deque_t;

void *xmalloc(size_t s) {
	void *p = malloc(s);
	if (p == NULL) {
		fprintf(stderr, "Can't malloc\n");
		exit(1);
	}
	return p;
}

deque_t *deque_init() {
	deque_t *d = xmalloc(sizeof(deque_t));
	d->head = xmalloc(sizeof(node_t));
	d->tail = xmalloc(sizeof(node_t));

	d->head->val = -1;
	d->head->prev = d->head;
	d->head->next = d->tail;

	d->tail->val = -2;
	d->tail->prev = d->head;
	d->tail->next = d->tail;

	d->fl = d->tail;
	return d;
}

node_t *node_get(deque_t *d) {
	node_t *p;
	if (d->fl == d->tail) {
		p = xmalloc(sizeof(node_t));
	} else {
		p = d->fl;
		d->fl = p->next;
	}
	return p;
}

void node_free(deque_t *d, node_t *p) {
	p->next->prev = p->prev;
	p->prev->next = p->next;
	p->next = d->fl;
	d->fl = p;
}

void insert(deque_t *d, int x) {
	node_t *p = node_get(d);
	p->val = x;
	p->next = d->head->next;
	p->prev = d->head;
	d->head->next->prev = p;
	d->head->next = p;
}

void delete(deque_t *d, int x) {
	node_t *p;
	for (p = d->head->next; p != d->tail; p = p->next) {
		if (p->val == x) {
			node_free(d, p);
			break;
		}
	}
}

void deleteFirst(deque_t *d) {
	if (d->head->next != d->tail) node_free(d, d->head->next);
}

void deleteLast(deque_t *d) {
	if (d->tail->prev != d->head) node_free(d, d->tail->prev);
}

void print(deque_t *d) {
	node_t *p;
	const char *sp = "";
	for (p = d->head->next; p != d->tail; p = p->next) {
		printf("%s%d", sp, p->val);
		sp = " ";
	}
	putchar('\n');
}

int main(int argc, char **argv)
{
	int n, i, x;
	char buf[128];
	deque_t *d;

	d = deque_init();
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%s", buf);
		if (strcmp(buf, "insert") == 0) {
			scanf("%d", &x);
			insert(d, x);
		} else if (strcmp(buf, "delete") == 0) {
			scanf("%d", &x);
			delete(d, x);
		} else if (strcmp(buf, "deleteFirst") == 0) {
			deleteFirst(d);
		} else if (strcmp(buf, "deleteLast") == 0) {
			deleteLast(d);
		} else {
			fprintf(stderr, "Command not found: %s\n", buf);
		}
	}

	print(d);

	return 0;
}

