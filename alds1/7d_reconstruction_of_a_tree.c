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

#define MAX_N 64

typedef struct node_st {
	int left, right;
	int pre, in;
} node_t;

int n;
int preorder[MAX_N];
int inorder[MAX_N];
node_t nd[MAX_N];

void recons(int id, int base , int n) {
	int pre, in, remain_left, remain_right;
	if (id == 0) return;

	pre = nd[id].pre;
	in = nd[id].in;
	remain_left = (in - base);
	remain_right = n - remain_left - 1;

	if (remain_left > 0) {
		nd[id].left = preorder[pre+1];
		recons(nd[id].left, base, remain_left);
	}

	if (remain_right > 0) {
		nd[id].right = preorder[pre+remain_left+1];
		recons(nd[id].right, in+1, remain_right);
	}
}

void postorder(int id) {
	static const char *sp = "";
	if (!id) return;
	if (nd[id].left) postorder(nd[id].left);
	if (nd[id].right) postorder(nd[id].right);
	printf("%s%d", sp, id);
	sp = " ";
}

int main(int argc, char **argv)
{
	int i;

	memset(nd, 0, sizeof(node_t) * MAX_N);
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf("%d", &preorder[i]);
		nd[preorder[i]].pre = i;
	}
	for (i = 0; i < n; i++) {
		scanf("%d", &inorder[i]);
		nd[inorder[i]].in = i;
	}

	recons(preorder[0], 0, n);
	postorder(preorder[0]);
	puts("");

	return 0;
}

