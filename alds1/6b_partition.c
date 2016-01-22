#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_N 100010

#define exchange(a, b) exchange_(&(a), &(b))
void exchange_(int *a, int *b) {
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

int Partition(int *A, int p, int r) {
	int x = A[r];
	int i = p - 1, j;
	for (j = p; j < r; j++) {
		if (A[j] <= x) {
			i++;
			exchange(A[i], A[j]);
		}
	}
	exchange(A[i+1], A[r]);
	return i+1;
}

int arr[MAX_N];
int n;

int main(int argc, char **argv)
{
	int i, p;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &arr[i]);
	}

	p = Partition(arr, 0, n-1);

	const char *fmt, *sp = "";
	for (i = 0; i < n; i++) {
		if (i == p) {
			fmt = "%s[%d]";
		} else {
			fmt = "%s%d";
		}
		printf(fmt, sp, arr[i]);
		sp = " ";
	}
	putchar('\n');
	return 0;
}

