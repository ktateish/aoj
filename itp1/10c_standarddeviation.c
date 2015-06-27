#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <math.h>
typedef long long ll;
const int MYINF = 1e9+7;

int n;
int s[1024];

double avg() {
	double ret = 0;
	int i;
	for (i = 0; i < n; i++) {
		ret += s[i];
	}
	return ret / n;
}

double sd() {
	int i;
	double m = avg(), ret = 0;

	for (i = 0; i < n; i++) {
		ret += (s[i] * 1.0 - m) * (s[i] * 1.0 - m);
	}

	return sqrt(ret / n);
}

int main(int argc, char **argv)
{
	int i;

	for (;;) {
		scanf("%d", &n);
		if (!n) break;
		for (i = 0; i < n; i++) {
			scanf("%d", &s[i]);
		}
		printf("%lf\n", sd());
	}
	return 0;
}

