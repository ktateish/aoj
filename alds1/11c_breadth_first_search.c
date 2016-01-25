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

#define MAX_N 128
#define MAX_Q (128*129)
int G[MAX_N][MAX_N];
int n;
int sz[MAX_N];
int d[MAX_N];
int visited[MAX_N];
int q[MAX_Q];
int qs, qe;

void enqueue(int x) {
	if ((qe + 1) % MAX_Q == qs % MAX_Q) exit(1);
	q[qe] = x;
	qe = (qe+1) % MAX_Q;
}

int dequeue() {
	if (qs == qe) return -1;
	int res = q[qs];
	qs = (qs+1) % MAX_Q;
	return res;
}

void bfs() {
	int v, i;

	for (i = 1; i <= n; i++) {
		visited[i] = false;
		d[i] = MYINF;
	}

	d[1] = 0;
	enqueue(1);
	while ((v = dequeue()) != -1) {
		if (visited[v]) continue;
		visited[v] = true;
		for (i = 0; i < sz[v]; i++) {
			int next = G[v][i];
			if (d[v] + 1 < d[next]) d[next] = d[v] + 1;
			if (!visited[next]) enqueue(next);
		}
	}
}

int main(int argc, char **argv)
{
	int i, j, u, k, v;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d %d", &u, &k);
		sz[u] = k;
		for (j = 0; j < sz[u]; j++) {
			scanf("%d", &v);
			G[u][j] = v;
		}
	}


	bfs();

	for (i = 1; i <= n; i++) {
		printf("%d %d\n", i, visited[i] ? d[i] : -1);
	}

	return 0;
}
