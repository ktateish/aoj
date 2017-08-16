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
 
int stk[256];
int top;
 
void push(int x) {
  stk[top++] = x;
}
 
int pop() {
  if (top == 0) {
    cerr << "stack is empty" << endl;
    exit(1);
  }
  return stk[--top];
}
 
int main(int argc, char **argv)
{
  vector<function<int(int,int)>> ops = {
    [](int a, int b){return a+b;},
    [](int a, int b){return a-b;},
    [](int a, int b){return a*b;},
  };
 
  string s;
  while (cin >> s) {
    auto i = string("+-*").find(s[0]);
    if (i != string::npos) {
      int b = pop();
      int a = pop();
      push(ops[i](a, b));
    } else {
      push(stoi(s));
    }
  }
  cout << pop() << endl;
  return 0;
}
