#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

const int B = 5;
const int F = 4;
const int R = 11;
static int arr[5][4][11];

int main(int argc, char **argv)
{
	int i, n, b, f, r, v;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d %d %d %d", &b, &f, &r, &v);
		arr[b][f][r] += v;
	}
	for (b = 1; b < B; b++) {
		for (f = 1; f < F; f++) {
			for (r = 1; r < R; r++) {
				printf(" %d", arr[b][f][r]);
			}
			putchar('\n');
		}
		if (b + 1 == B) break;
		for (i = 0; i < 20; i++) putchar('#');
		putchar('\n');
	}

	return 0;
}

