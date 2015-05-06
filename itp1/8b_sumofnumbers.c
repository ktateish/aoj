#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define len 1010
static char s[len];

int main(int argc, char **argv)
{
	int sum, i;
	for (;;) {
		fgets(s, len, stdin);
		if (s[0] == '0' && (s[1] == '\0' || s[1] == '\n')) break;
		sum = 0;
		for (i = 0; s[i] != '\0' && s[i] != '\n'; i++) {
			sum += s[i] - '0';
		}
		printf("%d\n", sum);
	}
	return 0;
}

