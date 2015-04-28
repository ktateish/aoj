#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int INF = 1e9+7;

int should_pirnt(int x) {
	int i;
	if (x % 3 == 0) return 1;
	else for (i = x; x > 0; x /= 10) {
		if (x % 10 == 3) return 1;
	}
	return 0;
}

void call(int n) {
	int i;
	for (i = 1; i <= n; i++) {
		if (should_pirnt(i)) printf(" %d", i);
	}
	putchar('\n');
}

int main(int argc, char **argv)
{
	int n;
	scanf("%d", &n);
	call(n);
	return 0;
}

