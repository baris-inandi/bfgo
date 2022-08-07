#include <iostream>
#include <unordered_map>
using namespace std;

class Bf
{
public:
    unordered_map<uint_fast16_t, uint_fast8_t> _t =
        unordered_map<uint_fast16_t, uint_fast8_t>();
    uint_fast16_t _p = 0;
    uint_fast8_t c() { return _t[_p]; }          // current
    bool w() { return c() != 0; }                // while, false if current is 0
    void p(uint_fast8_t x) { _t[_p] = c() + x; } // plus, increment
    void m(uint_fast8_t x) { _t[_p] = c() - x; } // minus, decrement
    void l(uint_fast8_t x) { _p -= x; }          // left
    void r(uint_fast8_t x) { _p += x; }          // right
    void o() { printf("%c", (char)c()); }        // out, print
    void i()
    {
        string x;
        getline(cin, x);
        for (uint_fast8_t i = 0; i < x.length(); i++)
        {
            _t[_p + i] = x[i];
        }
    } // in, read
};

void impl(Bf b) {}

int main()
{
    Bf b;
    impl(b);
    return 0;
};
