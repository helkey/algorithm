# 1115. Print FooBar Alternately

import threading
class FooBar:
    def __init__(self, n):
        self.n = n
        self.evFoo = threading.Event()
        self.evBar = threading.Event()

    def foo(self, printFoo: 'Callable[[], None]') -> None:
        for i in range(self.n):
            # printFoo() outputs "foo". Do not change or remove this line.
            printFoo()
            self.evBar.set()
            self.evFoo.wait()
            self.evFoo.clear()
            
    def bar(self, printBar: 'Callable[[], None]') -> None:
        for i in range(self.n):
            self.evBar.wait()
            self.evBar.clear()
            # printBar() outputs "bar". Do not change or remove this line.
            printBar()
            self.evFoo.set()
