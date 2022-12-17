import time
from os import system
from sys import argv

RUNS = 100
cmd = " ".join(argv[1:])

mean = 0.0
for i in range(RUNS):
    start = time.time()
    system(f"{cmd} >/dev/null")
    end = time.time()
    mean += (end - start) * 1000
mean /= RUNS

print(f"{RUNS} runs of \"{cmd}\": ", str(round(mean, 2)) + "ms")
