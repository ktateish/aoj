#include <stdio.h>
 
int parent(int i) {
    if (i) return (i - 1) /4;
    else return 0;
}
 
int child(int i, char c) {
    switch (c) {
    case 'A':
        return i*4 + 1;
        break;
    case 'C':
        return i*4 + 2;
        break;
    case 'G':
        return i*4 + 3;
        break;
    case 'T':
        return i*4 + 4;
        break;
    default:
        return -1;
    }
}
 
unsigned char H[2100000];
 
int index(int i) {
    return i / 8;
}
 
int offset(int i) {
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
 
void Insert(const char *s) {
    int nd = 0;
    const char *p;
    for (p = s; *p != '\0'; p++) {
        nd = child(nd, *p);
    }
    Set(nd);
}
 
int Find(const char *s) {
    int nd = 0;
    const char *p;
    for (p = s; *p != '\0'; p++) {
        nd = child(nd, *p);
    }
    return Get(nd);
}
 
int main(int argc, char **argv)
{
    int n, i;
    char s[32];
    scanf("%d", &n);
    for (i = 0; i < n; i++) {
        scanf("%s", s);
        if (s[0] == 'i') {
            scanf("%s", s);
            Insert(s);
        } else {
            scanf("%s", s);
            if (Find(s)) {
                printf("yes\n");
            } else {
                printf("no\n");
            }
        }
    }
 
    return 0;
}