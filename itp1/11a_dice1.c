#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;

typedef struct {
	int top;
	int north;
	int south;
	int east;
	int west;
	int bottom;
} dice;

void dice_init(dice *d, int v1, int v2, int v3, int v4, int v5, int v6) {
	d->top = v1;
	d->north = v5;
	d->south = v2;
	d->east = v3;
	d->west = v4;
	d->bottom = v6;
}

enum direction {
	DIR_N,
	DIR_S,
	DIR_E,
	DIR_W,
};

void rotate(dice *d, int dir) {
	int tmp;
	switch (dir) {
	case DIR_N:
		tmp = d->top;
		d->top = d->south;
		d->south = d->bottom;
		d->bottom = d->north;
		d->north = tmp;
		break;
	case DIR_S:
		tmp = d->top;
		d->top = d->north;
		d->north = d->bottom;
		d->bottom = d->south;
		d->south = tmp;
		break;
	case DIR_E:
		tmp = d->top;
		d->top = d->west;
		d->west = d->bottom;
		d->bottom = d->east;
		d->east = tmp;
		break;
	case DIR_W:
		tmp = d->top;
		d->top = d->east;
		d->east = d->bottom;
		d->bottom = d->west;
		d->west = tmp;
		break;
	default:
		fprintf(stderr, "Unspported direction: %d\n", dir);
		exit(1);
	}
	return;
}

int main(int argc, char **argv)
{
	int i, v[6], dir;
	dice d;
	char s[128];
	
	for (i = 0; i < 6; i++) {
		scanf("%d", &v[i]);
	}
	dice_init(&d, v[0], v[1], v[2], v[3], v[4], v[5]);
	scanf("%s", s);

	for (i = 0; i < strlen(s); i++) {
		dir = -1;
		switch (s[i]) {
		case 'N':
			dir = DIR_N;
			break;
		case 'S':
			dir = DIR_S;
			break;
		case 'E':
			dir = DIR_E;
			break;
		case 'W':
			dir = DIR_W;
			break;
		}
		rotate(&d, dir);
	}
	printf("%d\n", d.top);

	return 0;
}

