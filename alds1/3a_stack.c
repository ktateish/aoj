#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
typedef long long ll;
const int MYINF = 1e9+7;

ll S[256];
int idx;

void push(ll x) {
	S[idx++] = x;
}

ll pop() {
	if (idx == 0) {
		fprintf(stderr, "Stack is empty");
		exit(1);
	}
	return S[--idx];
}

int main(int argc, char **argv)
{
	ll x, y;
	char buf[16];

	while (scanf("%s", buf) > 0) {
		if (buf[0] == '+') {
			y = pop();
			x = pop();
			push(x + y);
		} else if (buf[0] == '-') {
			y = pop();
			x = pop();
			push(x - y);
		} else if (buf[0] == '*') {
			y = pop();
			x = pop();
			push(x * y);
		} else {
			push(atoll(buf));
		}
	}
	printf("%lld\n", pop());
	return 0;
}

