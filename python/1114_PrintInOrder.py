# 1114. Print in Order

TypeError: '_thread.lock' object is not callable

# Python offers mutexes, semaphores, and events for syncronization
import threading
import threading
class Foo:
    def __init__(self):
        self.lock2nd = threading.Lock()
        self.lock3rd = threading.Lock()
        self.lock2nd.acquire()
        self.lock3rd.acquire()

    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.lock2nd.release()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.lock2nd.acquire()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        # self.lock2nd.release()
        self.lock3rd.release()

    def third(self, printThird: 'Callable[[], None]') -> None:
        self.lock3rd.acquire()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()
        # self.lock3rd.release()

# USING EVENTS
class Foo:
    def __init__(self):
        self.ev2nd = threading.Event()
        self.ev3rd = threading.Event()

    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.ev2nd.set()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.ev2nd.wait()
        # printSecond() outputs "second". Do not change or remove this line.
        printSecond()
        self.ev3rd.set()
        
    def third(self, printThird: 'Callable[[], None]') -> None:
        self.ev3rd.wait()
        # printThird() outputs "third". Do not change or remove this line.
        printThird()

