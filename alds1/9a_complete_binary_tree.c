#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_H 264
#define parent(i) ((i)/2)
#define left(i) ((i)*2)
#define right(i) ((i)*2+1)

int H;
int heap[MAX_H];
#define len heap[0]

void insert(int x) {
	heap[++len] = x;
}

void print_node(int i) {
	printf("node %d: key = %d, ", i, heap[i]);
	if (0 < parent(i)) printf("parent key = %d, ", heap[parent(i)]);
	if (left(i) <= len) printf("left key = %d, ", heap[left(i)]);
	if (right(i) <= len) printf("right key = %d, ", heap[right(i)]);
	puts("");
}

int main(int argc, char **argv)
{
	int i, x;
	scanf("%d", &H);
	for (i = 0; i < H; i++) {
		scanf("%d", &x);
		insert(x);
	}

	for (i = 1; i <= H; i++) {
		print_node(i);
	}

	return 0;
}

