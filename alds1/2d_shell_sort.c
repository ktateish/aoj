#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
#define MYINF (1000007)

static int A[MYINF];
static int n;
static int cnt;
static int m;
static int G[128];

void insertion_sort(int g) {
	int i, j, v;
	for (i = g; i < n; i++) {
		v = A[i];
		j = i - g;
		while (j >= 0 && A[j] > v) {
			A[j+g] = A[j];
			j = j - g;
			cnt++;
		}
		A[j+g] = v;
	}
}

void shell_sort() {
	int i, j;

	cnt = 0;
	for (i = 0, j = n/2; i < 127 && j > 0; i++, j /= 2) {
		G[i] = j;
	}
	if (i == 0) {
		G[i] = 1;
		m = 1;
	} else {
		if (G[i-1] == 2) {
			G[i++] = 1;
		}
		m = i;
	}
	for (i = 0; i < m; i++) {
		insertion_sort(G[i]);
	}
}

int main(int argc, char **argv)
{
	int i;
	const char *sp = "";
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &A[i]);
	}
	shell_sort();
	printf("%d\n", m);
	for (i = 0; i < m; i++) {
		printf("%s%d", sp, G[i]);
		sp = " ";
	}
	putchar('\n');
	printf("%d\n", cnt);
	for (i = 0; i < n; i++) {
		printf("%d\n", A[i]);
	}

	return 0;
}

