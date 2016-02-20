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

int main(int argc, char **argv)
{
  int n, q;
  map<string, int> m;
  queue<string> waiting;
  cin >> n >> q;
  for (int i = 0; i < n; i++) {
    string name;
    int time;
    cin >> name >> time;
    m[name] = time;
    waiting.push(name);
  }

  int elapsed = 0;
  while (sz(waiting) > 0) {
    string s = waiting.front(); waiting.pop();
    m[s] -= q;
    if (m[s] > 0) {
      elapsed = elapsed + q;
      waiting.push(s);
    } else {
      elapsed = elapsed + q + m[s];
      cout << s << " " << elapsed << endl;
    }
  }
  return 0;
}

