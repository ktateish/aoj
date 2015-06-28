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

const long double EPS = 1e-5;

int main(int argc, char **argv)
{
  if (argc < 3) {
    fprintf(stderr, "Usage: %s <file1> <file2>\n", argv[0]);
    return 1;
  }

  ifstream in1(argv[1]);
  ifstream in2(argv[2]);

  char c1, c2;
  while (in1.get(c1)) {
    if (!in2.get(c2)) {
      fprintf(stderr, "%s is reached EOF, but %s isn't.", argv[2], argv[1]);
      return 1;
    }
    if (isspace(c1)) {
      if (c1 == c2) continue;
      fprintf(stderr, "%s and %s are differ. maybe a format error.", argv[1], argv[2]);
      return 1;
    } 
    in1.putback(c1);
    in2.putback(c2);

    string s1, s2;
    long double ldv1, ldv2;

    in1 >> s1;
    in2 >> s2;
    if (is_double(s1)) {
      if (!is_double(s2)) {
        fprintf(stderr, "%s and %s are differ. maybe a wrang answer.", argv[1], argv[2]);
        return 1;
      }
      ldv1 = stold(s1);
      ldv2 = stold(s2);
      if (abs(ldv1 - ldv2) < EPS) continue;
    } else {
      if (s1 == s2) continue;
    }
    fprintf(stderr, "%s and %s are differ. maybe a wrang answer.", argv[1], argv[2]);
    return 1;
  }
  if (in2.get(c2)) {
    fprintf(stderr, "%s is reached EOF, but %s isn't.", argv[1], argv[2]);
    return 1;
  }

  return 0;
}

