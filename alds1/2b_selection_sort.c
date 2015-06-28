#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

static int A[128];
int N, swap_count;

void swap(int i, int j) {
	int tmp;
	tmp = A[i];
	A[i] = A[j];
	A[j] = tmp;
	swap_count++;
}

void selection_sort() {
	int i, j, mini;
	for (i = 0; i < N; i++) {
		mini = i;
		for (j = i; j < N; j++) {
			if (A[j] < A[mini]) mini = j;
		}
		if (i != mini) swap(i, mini);
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
	selection_sort();
	for (i = 0; i < N; i++) {
		printf("%s%d", sp, A[i]);
		sp = " ";
	}
	putchar('\n');
	printf("%d\n", swap_count);
	return 0;
}

