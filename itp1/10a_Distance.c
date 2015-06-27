#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <math.h>
typedef long long ll;
const int MYINF = 1e9+7;

int main(int argc, char **argv)
{
	double x1, y1, x2, y2;

	scanf("%lf %lf %lf %lf", &x1, &y1, &x2, &y2);
	printf("%f\n", sqrt((x2 - x1) * (x2 - x1)
				+ (y2 - y1) * (y2 - y1)));
	return 0;
}

