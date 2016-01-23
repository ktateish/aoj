#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <math.h>
#include <complex.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

complex ARG60 = cos(M_PI/3)+sin(M_PI/3)*I;

void pprint(complex c) {
	printf("%lf %lf\n", creal(c), cimag(c));
}

void koch_curve_r(int n, complex p, complex q) {
	if (n == 0) {
		pprint(p);
		return;
	}
	complex pq = q - p;
	complex s = (pq/3);
	complex t = 2 * s;
	complex u = s * ARG60;
	koch_curve_r(n - 1, p, p+s);
	koch_curve_r(n - 1, p+s, p+s+u);
	koch_curve_r(n - 1, p+s+u, p+t);
	koch_curve_r(n - 1, p+t, q);
}

void koch_curve(int n) {
	complex s = 0+0*I;
	complex e = 100+0*I;
	koch_curve_r(n, s, e);
	pprint(e);
}

int main(int argc, char **argv)
{
	int n;
	scanf("%d", &n);
	koch_curve(n);
	return 0;
}

