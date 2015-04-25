#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int a, b;
	scanf("%d %d", &a, &b);
	printf("%d %d %.5f\n", a/b, a%b, (double)a / b);
	return 0;
}

