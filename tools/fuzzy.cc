#define dbg(fmt, ...) fprintf(stderr, fmt, __VA_ARGS__)
#define dpr(x) cerr<<#x<<": "<<x<<endl
#define dprc(c) do{cerr<<#c<<":";for(auto& _i:(c)){cerr<<" "<<_i;}cerr<<endl;}while(0)
#include <bits/stdc++.h>
//#include <cctype>
using namespace std;
typedef pair<int,int> pii;
typedef vector<int> vi;
typedef vector<vi> vvi;
typedef long long ll;
int INF = 1e9+7;
#define all(c) begin(c), end(c)
#define tr(c,i) for(auto i=begin(c);i!=end(c);i++)
#define rtr(c,i) for(auto i=(c).rbegin();i!=(c).rend();i++)
#define rep(i,n) for(auto i=0;i<(n);i++)
#define pb push_back
#define sz(c) int((c).size())

bool is_double(const string& s) {
  const char *cs = s.c_str();
  bool has_dot = false;

  for (int i = 0; i < sz(s); i++) {
    if (isdigit(cs[i])) continue;
    if (cs[i] == '.') {
      if (has_dot) return false;
      else has_dot = true;
      continue;
    }
    return false;
  }
  return has_dot;
}

int main(int argc, char **argv)
{
  char c;

  cout.precision(5);
  cout.setf(ios::floatfield, ios::fixed);

  while (cin.get(c)) {
    if (isspace(c)) {
      cout << c;
      continue;
    } 
    cin.putback(c);

    string s;
    long double ldv;

    cin >> s;
    if (is_double(s)) {
      ldv = stold(s);
      cout << ldv;
    } else {
      cout << s;
    }
    
  }

  return 0;
}

