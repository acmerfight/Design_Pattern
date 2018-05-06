import abc


class GirlFriend(object):

    def __init__(self, nationality, eyesColor, language):
        self.nationality = nationality
        self.eyesColor = eyesColor
        self.language = language


class Baby(object):

    def __init__(self, nationality):
        self.nationality = nationality


class AbstractFactory(object):
    """
    Declare an interface for operations that create abstract product
    objects.
    """
    __metaclass__ = abc.ABCMeta

    @abc.abstractmethod
    def create_myLove(self):
        pass

    @abc.abstractmethod
    def create_baby(self):
        pass


class IndianGirlFriendFactory(AbstractFactory):

    nationality = "Indian"
    eyesColor = "black"
    language = "Hindi"

    def create_myLove(self):
        return GirlFriend(self.nationality, self.eyesColor, self.language)

    def create_baby(self):
        return Baby(self.nationality)


class KoreanGirlFriendFactory(AbstractFactory):

    nationality = "Korean"
    eyesColor = "brown"
    language = "Korean"

    def create_myLove(self):
        return GirlFriend(self.nationality, self.eyesColor, self.language)

    def create_baby(self):
        return Baby(self.nationality)


factory = IndianGirlFriendFactory()
print(factory.create_myLove().eyesColor)
print(factory.create_baby().nationality)
factory = KoreanGirlFriendFactory()
print(factory.create_myLove().eyesColor)
print(factory.create_baby().nationality)

