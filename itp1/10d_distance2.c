#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <math.h>
typedef long long ll;
const int MYINF = 1e9+7;

int n;
int x[128];
int y[128];

double identity(double x) { return x; }

double distance(int degree) {
	int i;
	double tmp;
	double ret = 0;
	double (*rtfn)(double) = identity;
	if (degree == 1) ;
	else if (degree == 2) rtfn = sqrt;
	else if (degree == 3) rtfn = cbrt;
	else {
		for (i = 0; i < n; i++) {
			tmp = abs(x[i] - y[i]);
			if (tmp > ret) ret = tmp;
		}
		return ret;
	}
	for (i = 0; i < n; i++) {
		ret += pow(abs(x[i] - y[i]), degree);
	}
	return rtfn(ret);
}

int main(int argc, char **argv)
{
	int i;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &x[i]);
	}
	for (i = 0; i < n; i++) {
		scanf("%d", &y[i]);
	}
	printf("%lf\n", distance(1));
	printf("%lf\n", distance(2));
	printf("%lf\n", distance(3));
	printf("%lf\n", distance(-1));
	return 0;
}

