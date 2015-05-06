#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define slen 210
static char s[slen];
static char p[slen];

int main(int argc, char **argv)
{
	int i, len;

	fgets(s, slen, stdin);
	len = strlen(s) - 1;
	s[len] = '\0';
	for (i = 0; i < len; i++) {
		s[len+i] = s[i];
	}
	s[len+len] = '\0';

	fgets(p, slen, stdin);
	len = strlen(p) - 1;
	p[len] = '\0';

	if (strstr(s, p)) puts("Yes");
	else puts("No");

	return 0;
}

