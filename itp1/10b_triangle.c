#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <math.h>
typedef long long ll;
const int MYINF = 1e9+7;
const double EPS = 1e-6;

typedef struct {
	double x;
	double y;
} V;

V vector_sub(V *a, V *b) {
	V c;
	c.x = a->x - b->x;
	c.y = a->y - b->y;
	return c;
}

double vector_abs(V *a) {
	return sqrt(a->x*a->x + a->y*a->y);
}

double a, b, c, rad;

int main(int argc, char **argv)
{
	scanf("%lf %lf %lf", &a, &b, &c);
	V va, vb, vc;

	rad = c * M_PI / 180.0;
	va.x = a;
	va.y = 0;
	vb.x = b * cos(rad);
	vb.y = b * sin(rad);
	vc = vector_sub(&va, &vb);

	printf("%f\n", va.x * vb.y / 2.0);
	printf("%f\n", a + b + vector_abs(&vc));
	printf("%f\n", vb.y);
	return 0;
}

