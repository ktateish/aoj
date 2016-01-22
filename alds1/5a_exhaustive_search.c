#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

int n, q;
int A[32];

bool search(int key, int i, int sum) {
	if (sum == key) return true;
	if (i == n) return false;
	if (search(key, i + 1, sum + A[i])) return true;
	else return search(key, i + 1, sum);
}

int main(int argc, char **argv)
{
	int i, x;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &A[i]);
	}
	scanf("%d", &q);
	for (i = 0; i < q; i++) {
		scanf("%d", &x);
		if (search(x, 0, 0)) puts("yes");
		else puts("no");
	}
	return 0;
}

