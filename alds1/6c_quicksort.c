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

typedef struct {
	int suit;
	int rank;
	int origi_order;
} card_t;

#define exchange(a, b) exchange_(&(a), &(b))
void exchange_(card_t *a, card_t *b) {
	card_t tmp = *a;
	*a = *b;
	*b = tmp;
}

int Partition(card_t *A, int p, int r) {
	int x = A[r].rank;
	int i = p - 1, j;
	for (j = p; j < r; j++) {
		if (A[j].rank <= x) {
			i++;
			exchange(A[i], A[j]);
		}
	}
	exchange(A[i+1], A[r]);
	return i+1;
}

void Quicksort(card_t *A, int p, int r) {
	if (p < r) {
		int q = Partition(A, p, r);
		Quicksort(A, p, q-1);
		Quicksort(A, q+1, r);
	}
}

card_t arr[MAX_N];
int n;

int main(int argc, char **argv)
{
	int i, rank, stable;
	char suit[16];
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%s %d", suit, &rank);
		arr[i].suit = suit[0];
		arr[i].rank = rank;
		arr[i].origi_order = i;
	}

	Quicksort(arr, 0, n-1);

	stable = 1;
	for (i = 1; i < n; i++) {
		if (arr[i-1].rank != arr[i].rank) continue;
		if (arr[i-1].origi_order > arr[i].origi_order) {
			stable = 0;
			break;
		}
	}
	if (stable) puts("Stable");
	else puts("Not stable");

	for (i = 0; i < n; i++) {
		printf("%c %d\n", arr[i].suit, arr[i].rank);
	}
	return 0;
}

