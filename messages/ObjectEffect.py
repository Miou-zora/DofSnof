from CustomDataWrapper import CustomDataWrapper

class ObjectEffect:
    def __init__(self, actionId = 0) -> None:
        self.actionId = actionId

    def deserialize(self, input: CustomDataWrapper):
        self.actionId = input.readVarUhShort()
