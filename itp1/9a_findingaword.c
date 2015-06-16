#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
#include <ctype.h>
typedef long long ll;
const int INF = 1e9+7;

void stoupper(char *s) {
	char *p;
	for (p = s; *p != '\0'; p++) {
		*p = toupper(*p);
	}
}

int main(int argc, char **argv)
{
	int count;
	char W[16];
	char a[1024];

	count = 0;
	scanf("%s", W);
	stoupper(W);
	for (;;) {
		scanf("%s", a);
		if (strcmp(a, "END_OF_TEXT") == 0) break;
		stoupper(a);
		if (strcmp(a, W) == 0) count++;
	}
	printf("%d\n", count);
	return 0;
}

