#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;
typedef long long ll;

int main(int argc, char **argv)
{
	int i, n, a;
	ll min, max, sum;
	min = INF;
	max = -INF;
	sum = 0;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &a);
		if (a < min) min = a;
		if (a > max) max = a;
		sum += a;
	}
	printf("%lld %lld %lld\n", min, max, sum);
	return 0;
}

