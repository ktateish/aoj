#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;


void printA(const int *A, int N) {
	int i;
	char *sp = "";
	for (i = 0; i < N; i++) {
		printf("%s%d", sp, A[i]);
		sp = " ";
	}
	putchar('\n');
}

void insersion_sort(int *A, int N, void (*fn)(const int *, int)) {
	int i, j, key;
	if (fn != NULL) fn(A, N);
	for (i = 1; i < N; i++) {
		key = A[i];
		for (j = i - 1; j >= 0 && A[j] > key; j--) {
			A[j + 1] = A[j];
		}
		A[j + 1] = key;
		if (fn != NULL) fn(A, N);
	}
}

int main(int argc, char **argv)
{
	int N;
	int A[128];
	int i;
	scanf("%d", &N);
	for (i = 0; i < N; i++) {
		scanf("%d", &A[i]);
	}

	insersion_sort(A, N, printA);

	return 0;
}

