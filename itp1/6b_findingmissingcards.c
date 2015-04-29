#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <string.h>
typedef long long ll;
const int KTINF = 1e9+7;

enum suit_enum { S, H, C, D, };
static const char *suitname = "SHCD";
static int cards[4][14];

int main(int argc, char **argv)
{
	int c, i, j, n, suit, rank;
	scanf("%d", &n);
	for (i = 0; i < n; i++) {
		scanf(" %c", &c);
		switch (c) {
		case 'S': suit = S; break;
		case 'H': suit = H; break;
		case 'C': suit = C; break;
		case 'D': suit = D; break;
		default:
			printf("foo");
			return 2;
		}
		scanf(" %d", &rank);
		cards[suit][rank] = 1;
	}
	for (i = 0; i < 4; i++) {
		for (j = 1; j < 14; j++) {
			if (!cards[i][j]) printf("%c %d\n", suitname[i], j);
		}
	}
	return 0;
}

