#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int x, y;
	for (;;) {
		scanf("%d %d", &x, &y);
		if (x == 0 && y == 0) break;
		if (y < x) {
			x ^= y;
			y ^= x;
			x ^= y;
		}
		printf("%d %d\n", x, y);
	}
	return 0;
}

