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

void dice_copy(dice *d, const dice *s) {
	d->top = s->top;
	d->north = s->north;
	d->south = s->south;
	d->east = s->east;
	d->west = s->west;
	d->bottom = s->bottom;
}

enum direction {
	DIR_N,
	DIR_S,
	DIR_E,
	DIR_W,
	DIR_L,
	DIR_R,
};

void dice_rotate(dice *d, int dir) {
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

int match_all(const dice *d, const dice *s) {
	return (d->south == s->south)
		&& (d->east == s->east)
		&& (d->north == s->north)
		&& (d->west == s->west)
		&& (d->bottom == s->bottom);
}

int match_with_top(const dice *d, const dice *s) {
	dice nd;

	if (d->top != s->top) return 0;
	dice_copy(&nd, d);

	if (match_all(&nd, s)) return 1;
	dice_rotate(&nd, DIR_R);
	if (match_all(&nd, s)) return 1;
	dice_rotate(&nd, DIR_R);
	if (match_all(&nd, s)) return 1;
	dice_rotate(&nd, DIR_R);
	if (match_all(&nd, s)) return 1;

	return 0;
}

int match(const dice *d, const dice *s) {
	dice nd;
	dice_copy(&nd, d);
	
	if (match_with_top(&nd, s)) return 1;
	dice_rotate(&nd, DIR_N);
	if (match_with_top(&nd, s)) return 1;
	dice_rotate(&nd, DIR_N);
	if (match_with_top(&nd, s)) return 1;
	dice_rotate(&nd, DIR_N);
	if (match_with_top(&nd, s)) return 1;
	dice_rotate(&nd, DIR_E);
	if (match_with_top(&nd, s)) return 1;
	dice_rotate(&nd, DIR_W);
	dice_rotate(&nd, DIR_W);
	if (match_with_top(&nd, s)) return 1;

	return 0;
}

dice da[128];

int main(int argc, char **argv)
{
	int n, i, j, v[6];

	scanf("%d", &n);

	for (i = 0; i < n; i++) {
		for (j = 0; j < 6; j++) {
			scanf("%d", &v[j]);
		}
		dice_init(&da[i], v[0], v[1], v[2], v[3], v[4], v[5]);
	}

	for (i = 0; i < n - 1; i++) {
		for (j = i + 1; j < n; j++) {
			if (match(&da[i], &da[j])) {
				puts("No");
				return 0;
			}
		}
	}
	puts("Yes");
	return 0;
}

