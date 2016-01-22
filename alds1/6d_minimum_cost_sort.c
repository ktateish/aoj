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

#define MAX_N 1010

typedef struct {
	int w;
	int o;
} baggage;

int n;
int cost;
baggage W[MAX_N];
int visited[MAX_N];

int cmp(const void *a, const void *b) {
	return ((const baggage *)a)->w - ((const baggage *)b)->w;
}

void calc() {
	int i;
	qsort(W, n, sizeof(baggage), cmp);
	for (i = 0; i < n; i++) {
		if (visited[i]) continue;
		visited[i] = 1;

		int c1, c2, j;
		for (c1 = 0, j = W[i].o; j != i; j = W[j].o) {
			c1 += W[i].w + W[j].w;
			visited[j] = 1;
		}
		for (c2 = 2*(W[0].w+W[i].w), j=W[i].o; j != i; j = W[j].o) {
			c2 += W[0].w + W[j].w;
		}

		if (c1 < c2) cost += c1;
		else cost += c2;
	}
}

int main(int argc, char **argv)
{
	int i;

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &W[i].w);
		W[i].o = i;
	}
	calc();

	printf("%d\n", cost);
	return 0;
}

