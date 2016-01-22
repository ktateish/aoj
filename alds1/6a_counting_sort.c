#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_N (2000010)
#define MAX_K (10010)

int n, k;
int A[MAX_N];
int B[MAX_N];
int C[MAX_K];

void csort(void) {
	int i;
	for (i = 0; i < n; i++) {
		C[A[i]] = C[A[i]] + 1;
	}
	for (i = 1; i <= k; i++) {
		C[i] = C[i] + C[i - 1];
	}
	for (i = n-1; i >= 0; i--) {
		B[C[A[i]]] = A[i];
		C[A[i]] = C[A[i]] - 1;
	}
}

int main(int argc, char **argv)
{
	int i;
	scanf("%d", &n);
	k = 0;
	for (i = 0; i < n; i++) {
		scanf("%d", &A[i]);
		if (A[i] > k) k = A[i];
	}

	csort();
	const char *sp = "";
	for (i = 1; i <= n; i++) {
		printf("%s%d", sp, B[i]);
		sp = " ";
	}
	putchar('\n');
	return 0;
}

