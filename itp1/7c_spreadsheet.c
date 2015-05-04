#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

static int row[110];
static int sum[110];

int main(int argc, char **argv)
{
	int r, c, i, j;
	scanf("%d %d", &r, &c);
	for (i = 0; i < r; i++) {
		row[c] = 0;
		for (j = 0; j < c; j++) {
			scanf("%d", &row[j]);
			row[c] += row[j];
			sum[j] += row[j];
			printf("%d ", row[j]);
		}
		printf("%d\n", row[c]);
		sum[c] += row[c];
	}
	for (j = 0; j < c; j++) {
		printf("%d ", sum[j]);
	}
	printf("%d\n", sum[c]);
	return 0;
}

