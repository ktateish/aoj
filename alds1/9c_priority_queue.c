#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;
typedef int bool;
const bool true = 1;
const bool false = 0;

#define MAX_H 32000000
#define parent(i) ((i)/2)
#define left(i) ((i)*2)
#define right(i) ((i)*2+1)

int H;
int heap[MAX_H];
#define len heap[0]

void swap(int i, int j) {
	int t;
	t = heap[i];
	heap[i] = heap[j];
	heap[j] = t;
}

void shiftup(int i) {
	int p = parent(i);
	if (i == 1 || heap[p] > heap[i]) return;
	swap(p, i);
	shiftup(p);
}

void insert(int x) {
	heap[++len] = x;
	shiftup(len);
}

void shiftdown(int i) {
	int l = left(i), r = right(i), c = -1000;
	if (len < i) return;
	if (len < l) return;
	if (r <= len) {
		if (heap[l] > heap[i] && heap[r] > heap[i]) {
			if (heap[l] > heap[r]) c = l;
			else c = r;
		} else if (heap[l] > heap[i]) {
			c = l;
		} else if (heap[r] > heap[i]) {
			c = r;
		} else {
			return;
		}
	} else if (l <= len) {
		if (heap[l] > heap[i]) c = l;
		else return;
	}
	swap(c, i);
	shiftdown(c);
}

int extract(void) {
	int res;
	res = heap[1];
	heap[1] = heap[len--];
	shiftdown(1);
	return res;
}

int main(int argc, char **argv)
{
	int x;
	char cmd[256];

	while (1) {
		scanf("%s", cmd);
		if (!strcmp(cmd, "end")) break;
		if (!strcmp(cmd, "insert")) {
			scanf("%d", &x);
			insert(x);
		} else if (!strcmp(cmd, "extract")) {
			printf("%d\n", extract());
		}
	}

	return 0;
}

