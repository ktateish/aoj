#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

static int A[128];
int N, swap_count;

void swap(int *a, int *b) {
	int tmp;
	tmp = *a;
	*a = *b;
	*b = tmp;
	swap_count++;
}

void bubble_sort() {
	int i, j;
	for (i = 0; i < N; i++) {
		for (j = N - 1; j > i; j--) {
			if (A[j] < A[j-1]) swap(A+j, A+j-1);
		}
	}
}

int main(int argc, char **argv)
{
	int i;
	const char *sp = "";
	scanf("%d", &N);
	for (i = 0; i < N; i++) {
		scanf("%d", &A[i]);
	}
	bubble_sort();
	for (i = 0; i < N; i++) {
		printf("%s%d", sp, A[i]);
		sp = " ";
	}
	putchar('\n');
	printf("%d\n", swap_count);
	return 0;
}

