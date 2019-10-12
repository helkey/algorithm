# 1195. Fizz Buzz Multithreaded
# Faster than 53% of Python3 submissions

import threading
class FizzBuzz:
    def __init__(self, n: int):
        self.n = n
        self.sFizz = threading.Semaphore(0)
        self.sBuzz = threading.Semaphore(0)
        self.sFizzBuzz = threading.Semaphore(0)
        self.sNum = threading.Semaphore(1)
        
    # printNumber(x) outputs "x", where x is an integer.
    def number(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n + 1):
            self.sNum.acquire()
            self.fz = i % 3 == 0
            self.bz = i % 5 == 0
            if not self.fz and not self.bz:
                printNumber(i)
            self.sFizz.release()

    # printFizz() outputs "fizz"
    def fizz(self, printFizz: 'Callable[[], None]') -> None:
        for i in range(1, self.n + 1):
            self.sFizz.acquire()
            if self.fz and not self.bz:
                printFizz()
            self.sBuzz.release()
            
    # printBuzz() outputs "buzz"
    def buzz(self, printBuzz: 'Callable[[], None]') -> None:
        for i in range(1, self.n + 1):
            self.sBuzz.acquire()
            if self.bz and not self.fz:
                printBuzz()
            self.sFizzBuzz.release()

    # printFizzBuzz() outputs "fizzbuzz"
    def fizzbuzz(self, printFizzBuzz: 'Callable[[], None]') -> None:
        for i in range(1, self.n + 1):
            self.sFizzBuzz.acquire()
            if self.fz and self.bz:
                printFizzBuzz()
            self.sNum.release()

