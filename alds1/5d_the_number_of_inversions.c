#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define MAX_N (200010)
int n;
int A[MAX_N];
int S[MAX_N];

ll count;
void merge(int l, int r) {
	if (l >= r) return;
	merge(l, (l+r)/2);
	merge((l+r)/2+1, r);
	int i, j, k;
	for (k = 0, i = l, j = (l+r)/2+1; k < r - l + 1; k++) {
		if (i > (l+r)/2) {
			S[k] = A[j++];
		} else if (j > r) {
			S[k] = A[i++];
		} else {
			if (A[i] > A[j]) {
				S[k] = A[j++];
				count += (l+r)/2 - i + 1;
			} else {
				S[k] = A[i++];
			}
		}
	}
	for (i = 0; i < r - l + 1; i++) {
		A[l+i] = S[i];
	}
}

ll ninversions(void) {
	merge(0, n-1);
	return count;
}

int main(int argc, char **argv)
{
	int i;
	scanf("%d", &n);

	for (i = 0; i < n; i++) {
		scanf("%d", &A[i]);
		//S[i] = A[i];
	}
	printf("%lld\n", ninversions());


	return 0;
}

