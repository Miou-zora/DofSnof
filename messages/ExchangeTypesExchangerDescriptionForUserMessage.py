from CustomDataWrapper import CustomDataWrapper
all_items = {}
class ExchangeTypesExchangerDescriptionForUserMessage:
    def __init__(self) -> None:
        self.objectType = 0
        self.typeDescription = []
        self._isInitialized = True

    def deserialize(self, input: CustomDataWrapper):
        self._objectTypeFunc(input)
        _typeDescriptionLen = input.readUnsignedShort()
        for _ in range(_typeDescriptionLen):
            _val2 = input.readVarUhInt()
            if _val2 < 0:
                raise ValueError("Forbidden value (" + str(_val2) + ") on elements of typeDescription.")
            self.typeDescription.append(_val2)
    
    def _objectTypeFunc(self, input: CustomDataWrapper):
        self.objectType = input.readInt()
        if self.objectType < 0:
            raise ValueError("Forbidden value (" + str(self.objectType) + ") on element of ExchangeTypesExchangerDescriptionForUserMessage.objectType.")
        
    def __str__(self) -> str:
        # return f"objectType: {self.objectType}, typeDescription: [" + ', '.join([str(i) for i in self.typeDescription]) + "]"
        return f"objectType: {self.objectType}, typeDescription: [" + ', '.join([all_items[i] if i in all_items else str(i) for i in self.typeDescription]) + "]"
