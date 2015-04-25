#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int a, b, c, i, sum;
	scanf("%d %d %d", &a, &b, &c);
	for (i = a, sum = 0; i <= b && i <= c; i++) {
		if (c % i == 0) sum++;
	}
	printf("%d\n", sum);
	return 0;
}
