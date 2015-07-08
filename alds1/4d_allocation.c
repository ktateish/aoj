#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
bool true = 1;
bool false = 0;

int n, k;
int w[100010];

bool check(int capa) {
	int i, j, c;
	for (i = 0, j = 0; j < k; j++) {
		c = capa;
		while ((c - w[i]) >= 0) {
			c -= w[i];
			i++;
			if (i == n) {
				return true;
			}
		}
	}
	if (i == n) return true;
	return false;
}

int search(int l, int r) {
	int i, m;
	if (r - l < 5) {
		for (i = l - 1; i <= r + 1; i++) {
			if (check(i)) return i;
		}
	}
	m = (l + r) / 2;
	if (check(m)) {
		return search(l, m - 1);
	} else {
		return search(m + 1, r);
	}
}

int main(int argc, char **argv)
{
	int i, maxw, pd, d;
	scanf("%d %d", &n, &k);
	maxw = -1;
	for (i = 0; i < n; i++) {
		scanf("%d", &w[i]);
		if (maxw < w[i]) maxw = w[i];
	}

	if (k >= n) {
		printf("%d\n", maxw);
		return 0;
	}

	for (pd = d = maxw; !check(d); pd = d, d *= 2)
		;
	printf("%d\n", search(pd, d));

	return 0;
}

