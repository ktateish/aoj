#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

int A[500010];
int L[250010];
int R[250010];
int n, count;

bool cmp(int a, int b) {
	count++;
	return a <= b;
}

void merge(int left, int mid, int right) {
	int i, j, k, ni, nj;
	ni = mid - left;
	nj = right - mid;
	for (i = 0; i < ni; i++) {
		L[i] = A[left+i];
	}
	L[i] = MYINF;
	for (i = 0; i < nj; i++) {
		R[i] = A[mid+i];
	}
	R[i] = MYINF;
	for (i = 0, j = 0, k = left; k < right; k++) {
		if (cmp(L[i], R[j])) A[k] = L[i++];
		else A[k] = R[j++];
	}
}

void merge_sort(int left, int right) {
	int m;
	if (left + 1 == right) return;
	m = (left + right) / 2;
	merge_sort(left, m);
	merge_sort(m, right);
	merge(left, m, right);
}

int main(int argc, char **argv)
{
	int i;
	const char *sp;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &A[i]);
	}
	merge_sort(0, n);
	for (i = 0, sp = ""; i < n; i++) {
		printf("%s%d", sp, A[i]);
		sp = " ";
	}
	printf("\n%d\n", count);
	return 0;
}

