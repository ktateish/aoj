#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int i, x;
	for (i = 1; ; i++) {
		scanf("%d ", &x);
		if (x == 0) break;
		printf("Case %d: %d\n", i, x);
	}
	return 0;
}

