#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

int main(int argc, char **argv)
{
	int i, n, r, rmin, prof;

	scanf("%d", &n);
	scanf("%d", &rmin);
	scanf("%d", &r);
	prof = r - rmin;
	if (rmin > r) rmin = r;

	for (i = 2; i < n; i++) {
		scanf("%d", &r);
		if ((r - rmin) > prof) {
			prof = r - rmin;
		}
		if (r < rmin) rmin = r;
	}
	printf("%d\n", prof);

	return 0;
}

