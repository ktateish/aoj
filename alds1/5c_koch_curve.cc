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

void solve(vector<complex<double>>& v, int n,
      complex<double> p, complex<double> q) {
  if (n == 0) {
    v.pb(p);
    return;
  }
  complex<double> s = (q-p) * complex<double>(1.0/3, 0.0) + p;
  complex<double> t = (q-p) * complex<double>(2.0/3, 0.0) + p;
  double ang = M_PI/3;
  complex<double> u = (t-s) * complex<double>(cos(ang), sin(ang)) + s;
  solve(v, n-1, p, s);
  solve(v, n-1, s, u);
  solve(v, n-1, u, t);
  solve(v, n-1, t, q);
}

int main(int argc, char **argv)
{
  int n;
  scanf("%d", &n);
  vector<complex<double>> v;
  solve(v, n, complex<double>(0.0, 0.0), complex<double>(100.0, 0.0));
  for (auto& c:v) {
    printf("%f %f\n", real(c), imag(c));
  }
  printf("%f %f\n", 100.0, 0.0);

  return 0;
}

