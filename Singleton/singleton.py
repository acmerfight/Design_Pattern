# I don't really see the need, as a module with functions (and not a class) would serve well as a singleton. '
# All its variables would be bound to the module, which could not be instantiated repeatedly anyway.
# If you do wish to use a class, there is no way of creating private classes or private constructors in Python,
# so you can't protect against multiple instantiations, other than just via convention in use of your API.
# I would still just put methods in a module, and consider the module as the singleton.
# https://stackoverflow.com/questions/31875/is-there-a-simple-elegant-way-to-define-singletons


import threading


class SingletonMixin(object):
    __singleton_lock = threading.Lock()
    __singleton_instance = None

    @classmethod
    def instance(cls):
        if not cls.__singleton_instance:
            with cls.__singleton_lock:
                if not cls.__singleton_instance:
                    cls.__singleton_instance = cls()
        return cls.__singleton_instance
