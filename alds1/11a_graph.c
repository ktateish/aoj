#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_N 128

int G[MAX_N][MAX_N];
int n;

int main(int argc, char **argv)
{
	int i, j, u, k, v;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d %d", &u, &k);
		for (j = 0; j < k; j++) {
			scanf("%d", &v);
			G[u][v] = 1;
		}
	}

	for (i = 1; i <= n; i++) {
		const char *sp = "";
		for (j = 1; j <= n; j++) {
			printf("%s%d", sp, G[i][j]);
			sp = " ";
		}
		printf("\n");
	}

	return 0;
}

