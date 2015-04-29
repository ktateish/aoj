#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define MAX_N 110
#define MAX_M 110
static int a[MAX_N][MAX_M];
static int b[MAX_M];
static int c[MAX_N];

int main(int argc, char **argv)
{
	int i, j, n, m;
	scanf("%d %d", &n, &m);
	for (i = 0; i < n; i++) {
		for (j = 0; j < m; j++) {
			scanf("%d", &a[i][j]);
		}
	}
	for (i = 0; i < m; i++) {
		scanf("%d", &b[i]);
	}
	for (i = 0; i < n; i++) {
		for (j = 0; j < m; j++) {
			c[i] += a[i][j] * b[j];
		}
	}
	for (i = 0; i < n; i++) {
		printf("%d\n", c[i]);
	}
	return 0;
}

