class Color(object):

    def __init__(self, name):
        print("create")
        self.name = name


class ColorFactory(object):

    def __init__(self):
        self._instances = {}

    def create(self, name):
        # return self._instances.setdefault(name, Color(name))
        if name not in self._instances:
            self._instances[name] = Color(name)
        else:
            return self._instances[name]


colorFactory = ColorFactory()
colorFactory.create("red")
colorFactory.create("red")
