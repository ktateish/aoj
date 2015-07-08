#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define NSIZE 100010
#define QSIZE 50000

int S[NSIZE];
int n;

int cmp(const void *a, const void *b) {
	return *(int *)a - *(int *)b;
}

int binsearch_h(int x, int left, int right) {
	int m;
	if (left == right) {
		return S[left] == x;
	} else if (left > right) {
		return 0;
	}
	m = (right + left)  / 2;
	if (x < S[m]) {
		return binsearch_h(x, left, m - 1);
	} else if (x > S[m]) {
		return binsearch_h(x, m + 1, right);
	} else {
		return 1;
	}
}

int binsearch(int x) {
	return binsearch_h(x, 0, n - 1);
}

int main(int argc, char **argv)
{
	int q, i, x, count;

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &S[i]);
	}
	qsort(S, n, sizeof(int), cmp);

	scanf("%d", &q);
	count = 0;
	for (i = 0; i < q; i++) {
		scanf("%d", &x);
		count += binsearch(x);
	}
	printf("%d\n", count);

	return 0;
}

