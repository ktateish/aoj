#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int H, W, i, j;
	for (;;) {
		scanf("%d %d", &H, &W);
		if (H == 0 && W == 0) break;
		for (i = 0; i < H; i++) {
			for (j = 0; j < W; j++) {
				putchar('#');
			}
			putchar('\n');
		}
		putchar('\n');
	}

	return 0;
}

