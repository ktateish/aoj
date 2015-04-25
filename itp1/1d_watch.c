#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int h, m, s;
	scanf("%d", &s);
	m = s / 60;
	s = s % 60;
	h = m / 60;
	m = m % 60;
	printf("%d:%d:%d\n", h, m, s);
	return 0;
}

