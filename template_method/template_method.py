import abc


class AbstractClass(object):
    __metaclass__ = abc.ABCMeta

    def run(self):
        self._step_1()
        self._step_2()
        self._step_3()

    @staticmethod
    def _step_2():
        print("step2")

    @abc.abstractmethod
    def _step_1(self):
        pass

    @abc.abstractmethod
    def _step_3(self):
        pass


class ConcreteClass(AbstractClass):

    def _step_1(self):
        print("step1")

    def _step_3(self):
        print("step3")


if __name__ == "__main__":
    ConcreteClass().run()

