#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
typedef long long ll;
const int MYINF = 1e9+7;

const int nrtry = 30;

const int prtable[] = {
	2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61,
	67, 71, 73, 79, 83, 89, 97,
};
const int prtable_size = sizeof(prtable)/sizeof(int);

int gcd_rec(int x, int y) {
	if (y == 0) return x;
	return gcd_rec(y, x % y);
}

int gcd(int x, int y) {
	if (x < y) return gcd_rec(y, x);
	else return gcd_rec(x, y);
}

ll expmod(ll n, ll e, ll m) {
	ll nx;
	if (e == 1) {
		return n % m;
	} else if ((e % 2) == 0) {
		nx = expmod(n, e/2, m);
		return nx*nx % m;
	} else {
		nx = expmod(n, e - 1, m);
		return nx*(n%m) % m;
	}
}

int is_prime(int x) {
	int i, a;
	if (x < 2) return 0;
	for (i = 0; i < prtable_size && prtable[i] <= x; i++) {
		if (x == prtable[i]) return 1;
	}

	for (i = 0; i < nrtry; i++) {
try_again:
		a = random() % x;
		if (a < 3) goto try_again;
		if (gcd(a, x) != 1) return 0;
		if (expmod(a, x - 1, x) != 1LL) return 0;
	}

	return 1;
}

int main(int argc, char **argv)
{
	int i, n, x, count = 0;

	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &x);
		if (is_prime(x)) count++;
	}
	printf("%d\n", count);

	return 0;
}

