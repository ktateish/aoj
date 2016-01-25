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

#define MAX_N 1024
int q;
char X[MAX_N], Y[MAX_N];

int dp[MAX_N][MAX_N];
int xlen, ylen;

int max(int a, int b) {
	return a > b ? a : b;
}

int rec(int x, int y) {
	if (dp[x][y] >= 0) return dp[x][y];
	if (x == xlen || y == ylen) return dp[x][y] = 0;
	int m = max(rec(x, y+1), rec(x+1, y));
	if (X[x] == Y[y]) {
		return dp[x][y] = max(1+rec(x+1, y+1), m);
	} else {
		return dp[x][y] = m;
	}
}

int solve() {
	xlen = strlen(X);
	ylen = strlen(Y);
	memset(dp, -1, sizeof(dp));
	return rec(0, 0);
}

/*
int solve_by_loop() {
	xlen = strlen(X);
	ylen = strlen(Y);
	memset(dp, 0, sizeof(dp));
	int x, y;
	for (x = xlen - 1; x >= 0; x--) {
		for (y = ylen - 1; y >= 0; y--) {
			int m = max(dp[x][y+1], dp[x+1][y]);
			if (X[x] == Y[y]) dp[x][y] = max(1+dp[x+1][y+1], m);
			else dp[x][y] = m;
		}
	}
	return dp[0][0];
}
*/

int main(int argc, char **argv)
{
	int i;
	scanf("%d", &q);
	for (i = 0; i < q; i++) {
		scanf("%s %s", X, Y);
		printf("%d\n", solve());
	}

	return 0;
}

