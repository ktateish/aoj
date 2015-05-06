#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <ctype.h>
typedef long long ll;
const int MYINF = 1e9+7;

int main(int argc, char **argv)
{
	char *p = NULL;
	size_t i, n;
	getline(&p, &n, stdin);
	for (i = 0; i < n; i++) {
		if (p[i] == '\0') break;
		else if (islower(p[i])) putchar(toupper(p[i]));
		else if (isupper(p[i])) putchar(tolower(p[i]));
		else putchar(p[i]);
	}
	return 0;
}

