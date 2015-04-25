#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

void swap(int *x, int *y)
{
	int tmp = *x;
	*x = *y;
	*y = tmp;
}

int main(int argc, char **argv)
{
	int a, b, c;
	scanf("%d %d %d", &a, &b, &c);
	if (a > b) swap(&a, &b);
	if (b > c) swap(&b, &c);
	if (a > b) swap(&a, &b);
	printf("%d %d %d\n", a, b, c);
	return 0;
}

