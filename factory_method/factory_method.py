import abc


class ApplianceFactory(object):
    __metaclass__ = abc.ABCMeta

    type_name = ""

    @abc.abstractmethod
    def create(self, a_type):
        pass

    def start(self, a_type):
        appliance = self.create(a_type)
        print(appliance.upper())


class ChinaApplianceFactory(ApplianceFactory):

    def create(self, a_type):
        return "I am a ChinaApplianceFactory: " + a_type


class OtherApplianceFactory(ApplianceFactory):

    def create(self, a_type):
        return "I am a OtherApplianceFactory: " + a_type


if __name__ == "__main__":
    ChinaApplianceFactory().start("Stove")
    OtherApplianceFactory().start("Fridge")

