#define dbg(fmt,...) fprintf(stderr,fmt,__VA_ARGS__)
#define dpri(x) dbg(#x ": %d\n", x)
#define dprs(x) dbg(#x ": %s\n", x)
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
typedef long long ll;
const int MYINF = 1e9+7;

#define QUEUE_MAX 100010

const char *queue_n[QUEUE_MAX];
int queue_t[QUEUE_MAX];
int front, back;
int n;
int m;

int empty() {
	return front == back;
}

void enqueue(const char *name, int time) {
	if (back == ((front + 1) % m)) {
		fprintf(stderr, "Queue full\n");
		exit(1);
	}
	queue_n[front] = name;
	queue_t[front] = time;
	front = (front + 1) % m;
}

const char *dequeue(int *time) {
	const char *name;
	if (front == back) {
		fprintf(stderr, "Queue empty\n");
		exit(1);
	}
	name = queue_n[back];
	*time = queue_t[back];
	back = (back + 1) % m;
	return name;
}

int main(int argc, char **argv)
{
	int q, i, t, e, d;
	char s[16];
	const char *name;

	scanf("%d %d", &n, &q);
	m = n + 10;

	for (i = 0; i < n; i++) {
		scanf("%s %d", s, &t);
		enqueue(strdup(s), t);
	}

	e = 0;
	while (!empty()) {
		name = dequeue(&t);
		d = t - q;
		if (d > 0) {
			e += q;
			enqueue(name, d);
		} else {
			e += t;
			printf("%s %d\n", name, e);
		}
	}

	return 0;
}

