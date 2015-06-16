#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int INF = 1e9+7;

char str[1024];
char p[1024];

void cmd_p(int a, int b) {
	int i;
	for (i = a; i <= b; i++) {
		putchar(str[i]);
	}
	putchar('\n');
}

void cmd_rev(int a, int b) {
	int i;
	char c;
	for (i = a; i <= (a+b)/2; i++) {
		c = str[i];
		str[i] = str[b - i + a];
		str[b - i + a] = c;
	}
}

void cmd_rplc(int a, int b) {
	int i;

	scanf("%s", p);
	for (i = a; i <= b; i++) {
		str[i] = p[i - a];
	}
}

int main(int argc, char **argv)
{
	int q, a, b;
	char cmd[32];

	scanf("%s %d", str, &q);
	while (q--) {
		scanf("%s %d %d", cmd, &a, &b);
		if (strcmp(cmd, "print") == 0) {
			cmd_p(a, b);
		} else if (strcmp(cmd, "reverse") == 0) {
			cmd_rev(a, b);
		} else {
			cmd_rplc(a, b);
		}
	}
	return 0;
}

