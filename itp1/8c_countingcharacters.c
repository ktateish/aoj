#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <ctype.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define len 30
static int count[len];

int main(int argc, char **argv)
{
	int c, i;
	while ((c = getchar()) != EOF) {
		if (isalpha(c)) {
			count[tolower(c) - 'a']++;
		}
	}
	for (i = 0; i <= 'z' - 'a'; i++) {
		printf("%c : %d\n", 'a' + i, count[i]);
	}
	return 0;
}

