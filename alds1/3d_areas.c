#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define STKSZ 20480

typedef struct {
	int stk[STKSZ];
	int idx;
} stack_t;

void push(stack_t *s, int x) {
	s->stk[s->idx++] = x;
}

int pop(stack_t *s) {
	if (s->idx == 0) return -1;
	return s->stk[--s->idx];
}

int peek(stack_t *s) {
	if (s->idx == 0) return -1;
	return s->stk[s->idx - 1];
}

int empty(stack_t *s) {
	return s->idx == 0;
}

stack_t slope_stk;
stack_t lakes_left_stk;
stack_t lakes_capa_stk;
stack_t *slope = &slope_stk;
stack_t *lleft = &lakes_left_stk;
stack_t *lcapa = &lakes_capa_stk;

void merge(int newleft, int adding) {
	int l, c, sum;
	sum = adding;
	while (!empty(lleft)) {
		l = peek(lleft);
		c = peek(lcapa);
		if (newleft < l) {
			sum += c;
			pop(lleft);
			pop(lcapa);
		} else break;
	}
	push(lleft, newleft);
	push(lcapa, sum);
}

int main(int argc, char **argv)
{
	int c, xpos, left, potion, k, sum;

	xpos = 0;
	while ((c = getchar()) != EOF) {
		if (c == '\\') {
			push(slope, xpos);
		} else if (c == '/') {
			left = pop(slope);
			if (left == -1) continue;
			potion = xpos - left;
			merge(left, potion);
		}
		xpos++;
	}

	while (!empty(lleft)) pop(lleft); // for reuse
	sum = k = 0;
	while (!empty(lcapa)) {
		k++;
		c = pop(lcapa);
		sum += c;
		push(lleft, c);
	}
	printf("%d\n%d", sum, k);
	while (!empty(lleft)) {
		printf(" %d", pop(lleft));
	}
	putchar('\n');

	return 0;
}

