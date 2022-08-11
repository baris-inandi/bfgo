package compiler

const IR string = `#include <stdio.h>
int main()
{
    int t[30000] = {0};
    int *p = t;
    // ir %s
    return 0;
}
`
