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
	DIR_L,
	DIR_R,
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
	case DIR_L:
		tmp = d->north;
		d->north = d->east;
		d->east = d->south;
		d->south = d->west;
		d->west = tmp;
	case DIR_R:
		tmp = d->north;
		d->north = d->west;
		d->west = d->south;
		d->south = d->east;
		d->east = tmp;
		break;
	default:
		fprintf(stderr, "Unspported direction: %d\n", dir);
		exit(1);
	}
	return;
}

void dice_dump(dice *d) {
	printf("top:    %d\n", d->top);
	printf("east:   %d\n", d->east);
	printf("bottom: %d\n", d->bottom);
	printf("west:   %d\n", d->west);
	printf("north:  %d\n", d->north);
	printf("south:  %d\n", d->south);
}

int match_top(dice *d, const dice*to) {
	int i; 
	if (d->top == to->top) return 1;
	for (i = 0; i < 3; i++) {
		rotate(d, DIR_E);
		if (d->top == to->top) return 1;
	}
	if (d->north == to->top) {
		rotate(d, DIR_S);
		return 1;
	}
	if (d->south == to->top) {
		rotate(d, DIR_N);
		return 1;
	}
	return 0;
}

int match_top_south(dice *d, const dice *to) {
	int i;
	if (match_top(d, to)) {
		if (d->south == to->south) return 1;
		for (i = 0; i < 3; i++) {
			rotate(d, DIR_R);
			if (d->south == to->south) return 1;
		}
	}
	return 0;
}

int main(int argc, char **argv)
{
	int i, v[6], q, top, south;
	dice d, e;
	
	for (i = 0; i < 6; i++) {
		scanf("%d", &v[i]);
	}
	dice_init(&d, v[0], v[1], v[2], v[3], v[4], v[5]);

	scanf("%d", &q);
	for (i = 0; i < q; i++) {
		scanf("%d %d", &top, &south);
		dice_init(&e, top, south, -1, -1, -1, -1);
		if (match_top_south(&d, &e)) {
			printf("%d\n", d.east);
		}
		else printf("-1");
	}
	return 0;
}

