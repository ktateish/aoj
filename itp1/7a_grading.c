#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

int main(int argc, char **argv)
{
	int m, f, r;

	for (;;) {
		int mf;
		char res = 'F';
		scanf("%d %d %d", &m, &f, &r);
		mf = m + f;
		if (m < 0 && f < 0 && r < 0) break;
		else if (m < 0 || f < 0) res = 'F';
		else if (mf >= 80) res = 'A';
		else if (mf >= 65) res = 'B';
		else if (mf >= 50 || r >= 50) res = 'C';
		else if (mf >= 30) res = 'D';

		printf("%c\n", res);
	}
	return 0;
}

