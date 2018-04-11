import math
import time


def decorator(func):
    def timeit():
        start = time.time()
        func()
        print(time.time() - start)
    return timeit


@decorator
def mypow():
    math.pow(2, 100)


mypow()
