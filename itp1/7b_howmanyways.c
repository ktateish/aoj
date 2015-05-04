#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

int main(int argc, char **argv)
{
	int n, x, i, j, k, sum;
	for (;;) {
		sum = 0;
		scanf("%d %d", &n, &x);
		if (n == 0 && x == 0) break;
		for (i = 1; i <= n - 2; i++) {
			for (j = i+1; j <= n - 1; j++) {
				for (k = j+1; k <= n; k++) {
					if (i + j + k == x) sum++;
				}
			}
		}
		printf("%d\n", sum);
	}
	return 0;
}

