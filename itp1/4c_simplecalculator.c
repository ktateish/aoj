#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
const int INF = 1e9+7;

int main(int argc, char **argv)
{
	int a, b, res;
	char op;
	for (;;) {
		scanf("%d %c %d", &a, &op, &b);
		switch (op) {
		case '+': res = a + b; break;
		case '-': res = a - b; break;
		case '*': res = a * b; break;
		case '/': res = a / b; break;
		default: return 0;
		}
		printf("%d\n", res);
	}
	return 0;
}

