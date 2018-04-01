import abc


class Subject(object):
    def __init__(self):
        self._observers = set()

    def attach(self, observer):
        self._observers.add(observer)

    def detach(self, observer):
        self._observers.remove(observer)

    def notify(self, data):
        for observer in self._observers:
            observer.update(data)


class Observer(object):
    __metaclass__ = abc.ABCMeta

    @abc.abstractmethod
    def update(self, data):
        pass


class Observer1(Observer):

    def update(self, data):
        print(self.__class__.__name__ + ": " + data)


class Observer2(Observer):

    def update(self, data):
        print(self.__class__.__name__ + ": " + data)


if __name__ == '__main__':
    subject = Subject()
    observer1 = Observer1()
    observer2 = Observer2()
    subject.attach(observer1)
    subject.attach(observer2)
    subject.notify("hello")

    print("remove observer2")

    subject.detach(observer2)
    subject.notify("hello")



