#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_N 110
int n;
int rcs[MAX_N];

#define ROW(i) (rcs[i])
#define COL(i) (rcs[i+1])

int dp[MAX_N][MAX_N];

int min(int a, int b) {
	if (a < b) return a;
	else return b;
}

/* returns minimum value of i, i+1, i+2, ..., i+size matrix multiplication */
int rec(int i, int size) {
	if (size < 2) {
		return 0;
	} else if (size == 2) {
		return dp[i][size] = ROW(i)*COL(i)*COL(i+1);
	} else if (dp[i][size]) {
		return dp[i][size];
	}

	int j, m = MYINF;
	for (j = 1; j < size; j++) {
		m = min(m,   rec(i, j)
			   + rec(i+j, size-j)
			   + ROW(i)*ROW(i+j)*COL(i+size-1));
	}
	return dp[i][size] = m;
}

int solve() {
	return rec(0, n);
}

int main(int argc, char **argv)
{
	int i;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d %d", &rcs[i], &rcs[i+1]);
	}

	printf("%d\n", solve());

	return 0;
}

