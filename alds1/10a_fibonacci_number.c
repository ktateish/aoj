#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

int a[50];
int fib (int n) {
	if (a[n]) return a[n];
	if (n < 2) return a[n] = 1;
	return a[n] = fib(n-1) + fib(n-2);
}

int main(int argc, char **argv)
{
	int n;
	scanf("%d", &n);
	printf("%d\n", fib(n));
	return 0;
}

