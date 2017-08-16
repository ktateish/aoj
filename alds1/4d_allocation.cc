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

int n, k;
vi w;

bool check(int p) {
  int l = 0, q = 0;
  for (int i = 0; i < n; i++) {
    if (q < w[i]) {
      q = p;
      l++;
      if (l > k || q < w[i]) return false;
    }
    q -= w[i];
  }
  return true;
}

int search_minp(int minp, int maxp) {
  if (maxp <= minp) {
    for (int i = minp; i < maxp+2; i++) {
      if (check(i)) return i;
    }
    assert(0);
  }
  int med = (minp+maxp)/2;

  if (check(med)) return search_minp(minp, med-1);
  else return search_minp(med+1, maxp);
}

int main(int argc, char **argv)
{
  scanf("%d %d", &n, &k);
  w = vi(n);
  int maxp = -INF, total = 0;
  for (int i = 0; i < n; i++) {
    scanf("%d", &w[i]);
    maxp = max(maxp, w[i]);
    total += w[i];
  }
  printf("%d\n", search_minp(maxp-1, total));
  return 0;
}

