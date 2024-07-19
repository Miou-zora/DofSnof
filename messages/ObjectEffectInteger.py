from .ObjectEffect import ObjectEffect
from CustomDataWrapper import CustomDataWrapper

class ObjectEffectInteger(ObjectEffect):
    def __init__(self, actionId = 0, value = 0) -> None:
        super().__init__(actionId)
        self.value = value

    def deserialize(self, input: CustomDataWrapper):
        super().deserialize(input)
        self.value = input.readVarUhInt()
        if self.value < 0:
            raise ValueError("Forbidden value (" + str(self.value) + ") on element of ObjectEffectInteger.value.")
        
    def __str__(self) -> str:
        return f"actionId: {self.actionId}, value: {self.value}"
