#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_H 500010
#define parent(i) ((i)/2)
#define left(i) ((i)*2)
#define right(i) ((i)*2+1)

int H;
int heap[MAX_H];
#define len heap[0]

void maxHeapify(int i) {
	int l, r, *A, largest;
	A = heap;
	l = left(i);
	r = right(i);

	if (l <= H && A[l] > A[i]) largest = l;
	else largest = i;

	if (r <= H && A[r] > A[largest]) largest = r;

	if (largest != i) {
		int t;
		t = A[i];
		A[i] = A[largest];
		A[largest] = t;
		maxHeapify(largest);
	}
}

void buildMaxHeap() {
	int i;
	for (i = H/2; 0 < i; i--) {
		maxHeapify(i);
	}
}

void insert(int x) {
	heap[++len] = x;
}

//void insert(int x) {
//	int i, t;
//	heap[++len] = x;
//	for (i = len; 0 < parent(i) && heap[parent(i)] < heap[i]; i = parent(i)) {
//		t = heap[parent(i)];
//		heap[parent(i)] = heap[i];
//		heap[i] = t;
//	}
//}

int main(int argc, char **argv)
{
	int i, x;
	scanf("%d", &H);
	for (i = 0; i < H; i++) {
		scanf("%d", &x);
		insert(x);
	}
	buildMaxHeap();

	for (i = 1; i <= H; i++) {
		printf(" %d", heap[i]);
	}
	puts("");

	return 0;
}

