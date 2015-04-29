#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int KTINF = 1e9+7;

int main(int argc, char **argv)
{
	int n, a[110], i;
	char *sp = "";
	scanf("%d", &n);
	if (n == 0) return 0;
	for (i = 0; i < n; i++) {
		scanf("%d", a + i);
	}
	for (i = 0; i < n; i++) {
		printf("%s%d", sp, a[n - i - 1]);
		sp = " ";
	}
	putchar('\n');

	return 0;
}

