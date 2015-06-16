#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int n, taro, hanako, res;
	char tcard[128], hcard[128];

	taro = hanako = 0;
	scanf("%d", &n);
	while (n--) {
		scanf("%s %s", tcard, hcard);
		res = strcmp(tcard, hcard);
		if (res > 0) {
			taro += 3;
		} else if (res < 0) {
			hanako += 3;
		} else {
			taro += 1;
			hanako += 1;
		}
	}
	printf("%d %d\n", taro, hanako);
	return 0;
}

