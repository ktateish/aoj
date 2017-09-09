#include <stdio.h>

int parent(int i) {
	return (i - 1) /4;
}

int childTab[256];

void initChildTab() {
	childTab['A'] = 1;
	childTab['C'] = 2;
	childTab['G'] = 3;
	childTab['T'] = 4;
}

inline int child(int i, int c) {
	return i*4 + childTab[c];
}

unsigned char H[2100000];

inline int index(int i) {
	return i / 8;
}

inline int offset(int i) {
	return i % 8;
}

void Set(int i) {
	int j = index(i);
	int off = offset(i);
	H[j] = H[j] | (1 << off);
}

int Get(int i) {
	int j = index(i);
	int off = offset(i);
	return H[j] & (1 << off);
}

char buf[BUFSIZ];

int main(int argc, char **argv)
{
	int n, i, j, nd;
	char s[32];
	initChildTab();
	setbuffer(stdout, buf, BUFSIZ);
	scanf("%d", &n);
	for (j = 0; j < n; j++) {
		scanf("%s", s);
		if (s[0] == 'i') {
			scanf("%s", s);
			nd = 0;
			for (i = 0; s[i] != '\0'; i++) {
				nd = child(nd, s[i]);
			}
			Set(nd);
		} else {
			scanf("%s", s);
			nd = 0;
			for (i = 0; s[i] != '\0'; i++) {
				nd = child(nd, s[i]);
			}
			if (Get(nd)) {
				printf("yes\n");
			} else {
				printf("no\n");
			}
		}
	}
	fflush(stdout);

	return 0;
}

