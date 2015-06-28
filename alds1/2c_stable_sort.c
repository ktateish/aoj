#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
typedef long long ll;
const int MYINF = 1e9+7;

static char save_suit[64];
static int  save_rank[64];
static int  save_seq[64];
static char suit[64];
static int  rank[64];
static int  seq[64];
static int N;

void swap(int i, int j) {
	int tmp;
	tmp = suit[i];
	suit[i] = suit[j];
	suit[j] = tmp;

	tmp = rank[i];
	rank[i] = rank[j];
	rank[j] = tmp;

	tmp = seq[i];
	seq[i] = seq[j];
	seq[j] = tmp;
}

void bubble_sort() {
	int i, j;
	for (i = 0; i < N; i++) {
		for (j = N - 1; j > i; j--) {
			if (rank[j] < rank[j-1]) swap(j, j-1);
		}
	}
}

void selection_sort() {
	int i, j, mini;
	for (i = 0; i < N; i++) {
		mini = i;
		for (j = i; j < N; j++) {
			if (rank[j] < rank[mini]) mini = j;
		}
		if (i != mini) swap(i, mini);
	}
}

int is_stable() {
	int i, prev;
	prev = seq[0];
	for (i = 1; i < N; i++) {
		if (rank[i] > rank[i-1]) {
			prev = seq[i];
			continue;
		}
		if (seq[i] < seq[i-1]) return 0;
	}
	return 1;
}

void reset() {
	int i;
	for (i = 0; i < N; i++) {
		suit[i] = save_suit[i];
		rank[i] = save_rank[i];
		seq[i] = save_seq[i];
	}
}

void print_cards() {
	int i;
	const char *sp = "";
	for (i = 0; i < N; i++) {
		printf("%s%c%d", sp, suit[i], rank[i]);
		sp = " ";
	}
	putchar('\n');
}

int main(int argc, char **argv)
{
	int i;
	char buf[16];

	scanf("%d", &N);
	for (i = 0; i < N; i++) {
		scanf("%s", buf);
		save_suit[i] = buf[0];
		save_rank[i] = buf[1] - '0';
		save_seq[i] = i;
	}

	reset();
	bubble_sort();
	print_cards();
	if (is_stable()) puts("Stable");
	else puts("Not stable");

	reset();
	selection_sort();
	print_cards();
	if (is_stable()) puts("Stable");
	else puts("Not stable");

	return 0;
}

