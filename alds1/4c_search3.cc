#define dbg(...) fprintf(stderr,__VA_ARGS__)
#define dpr(x) cerr<<#x<<": "<<x<<endl;
#define dprc(c) do{cerr<<#c<<":";for(auto&_i:(c)){cerr<<" "<<_i;}cerr<<endl;}while(0)
#include <bits/stdc++.h>
using namespace std;
typedef pair<int, int> pii;
typedef vector<int> vi;
typedef vector<vi> vvi;
int INF = 1e9+7;
#define all(c) (c).begin(), (c).end()
#define tr(i,c) for(auto i=(c).begin();i!=(c).end();i++)
#define rtr(i,c) for(auto i=(c).rbegin();i!=(c).rend();i++)
#define rep(i,b) for(auto i=0;i<(b);i++)
#define pb push_back
#define sz(c) int((c).size())

#define HASH_SIZE (1000003)

enum ALPHABET {
  C,
  T,
  G,
  A,
  NUL,
};

struct node {
  string key;
  struct node *next;
};

struct node *H[HASH_SIZE];
struct node H_instance[HASH_SIZE];

int calchash(string& s) {
  int res = 0;
  for (int i = 0; i < sz(s); i++) {
    res = (res*47 + s[i])%HASH_SIZE;
  }
  return res;
}

void insert(string& s) {
  int h = calchash(s);

  node *p = H[h];
  while (p->next != nullptr) {
    if (p->key == s) return;
    p = p->next;
  }
  node *n = new node;
  n->key = s;
  n->next = nullptr;
  p->next = n;
}

bool find(string& s) {
  int h = calchash(s);

  for (node *p = H[h]->next; p != nullptr; p = p->next) {
    if (p->key == s) return true;
  }
  return false;
}

int main(int argc, char **argv)
{
  for (int i = 0; i < HASH_SIZE; i++) {
    H[i] = &H_instance[i];
    H[i]->key = string("x"); // will not be matched with input
    H[i]->next = nullptr;
  }

  int n;
  scanf("%d", &n);
  for (int i = 0; i < n; i++) {
    char cmd[16], sc[16];
    scanf("%s %s", cmd, sc);
    string s(sc);
    if (cmd[0] == 'i') insert(s);
    else if (cmd[0] == 'f') {
     if (find(s)) cout << "yes" << endl;
     else cout << "no" << endl;
    }
  }
  return 0;
}

