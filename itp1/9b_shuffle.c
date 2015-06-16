#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int INF = 1e9+7;

void reverse(char *s, int n) {
	int i;
	char c;
	for (i = 0; i < n/2; i++) {
		c = s[i];
		s[i] = s[n - i - 1];
		s[n - i - 1] = c;
	}
}

void rotate(char *s, int n) {
	int len = strlen(s);
	reverse(s, n);
	reverse(s+n, len-n);
	reverse(s, len);
}

int main(int argc, char **argv)
{
	int i, m, h;
	char deck[256];
	
	for (;;) {
		scanf("%s", deck);
		if (strcmp(deck, "-") == 0) break;
		scanf("%d", &m);
		i = strlen(deck);
		for (i = 0; i < m; i++) {
			scanf("%d", &h);
			rotate(deck, h);
		}
		printf("%s\n", deck);
	}
	return 0;
}

