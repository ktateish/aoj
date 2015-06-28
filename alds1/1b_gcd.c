#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

int gcd_rec(int x, int y) {
	if (y == 0) return x;
	return gcd_rec(y, x % y);
}

int gcd(int x, int y) {
	if (x < y) return gcd_rec(y, x);
	else return gcd_rec(x, y);
}

int main(int argc, char **argv)
{
	int a, b;

	scanf("%d %d", &a, &b);
	printf("%d\n", gcd(a, b));

	return 0;
}

