package boilerplate

const C_IR_BOILERPLATE string = `#include <stdio.h>
int main()
{
    %s t[%d] = {%d};
    %s *p = t;
    // ir %s
    return 0;
}
`
