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

struct lake {
  int pos;
  int capa;
};

int A;
stack<int> left_slope;
stack<lake> lakes;

void merge_lakes(lake& lk) {
}

int main(int argc, char **argv)
{
  int c, cpos;

  cpos = 0;
  while ((c = getchar()) != EOF) {
    if (c == '\\') left_slope.push(cpos);
    else if (c == '/' && sz(left_slope)) {  // 左側に斜面がなければ湖はつくれない
      int l = left_slope.top(); left_slope.pop();
      lake lk = {
        l,
        cpos - l, // 新しくわかった表層部分(0.5 + (cpos - left - 1) + 0.5)
      };
      A += lk.capa;
      // l より 右側にある湖は全部 lk に含まれる
      while (sz(lakes)) {
        lake r = lakes.top(); lakes.pop();
        if (r.pos < lk.pos) {
          lakes.push(r);
          break;
        }
        lk.capa += r.capa;
      }
      lakes.push(lk);
    }
    cpos++;
  }

  vector<int> v;
  while (sz(lakes)) {
    lake lk = lakes.top(); lakes.pop();
    v.pb(lk.capa);
  }
  reverse(all(v));
  printf("%d\n", A);
  printf("%d", sz(v));
  for (auto& x:v) {
    printf(" %d", x);
  }
  putchar('\n');
  return 0;
}

