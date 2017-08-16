#define dbg(...) fprintf(stderr,__VA_ARGS__)
#define dpr(x) cerr<<#x<<": "<<x<<endl;
#define dprc(c) do{cerr<<#c<<":";for(auto&_i:(c)){cerr<<" "<<_i;}cerr<<endl;}while(0)
#include <bits/stdc++.h>
using namespace std;
typedef pair<int, int> pii;
typedef vector<int> vi;
typedef vector<vi> vvi;
int INF = 1e9+7;
#define all(c) begin(c), end(c)
#define tr(i,c) for(auto i=begin(c);i!=end(c);i++)
#define rtr(i,c) for(auto i=(c).rbegin();i!=(c).rend();i++)
#define rep(i,b) for(auto i=0;i<(b);i++)
#define pb push_back
#define sz(c) int((c).size())

int main(int argc, char **argv)
{
  int n = 1000000;
  int m = 1000000;
  int q = 10000;
  //int n = 100;
  //int m = 100;
  //int q = 10;

  vvi E(n);
  set<pii> memo;
  vector<pii> Q;

  for (int i = 0; i < m; i++) {
tryagain:
    int x = random() % n;
    int y = random() % n;
    if (x != y && memo.find(pii(x, y)) == memo.end()) {
      E[x].pb(y);
      E[y].pb(x);
      memo.insert(pii(x, y));
    } else {
      goto tryagain;
    }
  }

  memo = set<pii>();
  for (int i = 0; i < q; i++) {
tryagain2:
    int x = random() % n;
    int y = random() % n;
    if (memo.find(pii(x, y)) == memo.end()) {
      Q.pb(pii(x, y));
      memo.insert(pii(x, y));
    } else {
      goto tryagain2;
    }
  }

  cout << n << " " << m << endl;
  for (int i = 0; i < n; i++) {
    for (auto x:E[i]) {
      if (i < x) cout << i << " " << x << endl;
    }
  }
  cout << q << endl;
  for (auto& p:Q) {
    cout << p.first << " " << p.second << endl;
  }

  return 0;
}

