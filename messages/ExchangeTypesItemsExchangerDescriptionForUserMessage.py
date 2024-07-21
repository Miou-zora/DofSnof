from CustomDataWrapper import CustomDataWrapper
from .BidExchangerObjectInfo import BidExchangerObjectInfo

all_items = {}

class ExchangeTypesItemsExchangerDescriptionForUserMessage:
    def __init__(self, objectGID = 0, objectType = 0, itemTypeDescriptions = None):
        self.objectGID = objectGID
        self.objectType = objectType
        self.itemTypeDescriptions = []
        self._isInitialized = True

    def getMessageId(self) -> int:
        return 2738

    def deserialize(self, input: CustomDataWrapper):
        self._objectGIDFunc(input)
        self._objectTypeFunc(input)
        _itemTypeDescriptionsLen: int = input.readUnsignedShort()
        for _ in range(_itemTypeDescriptionsLen):
            _item3 = BidExchangerObjectInfo()
            _item3.deserialize(input)
            self.itemTypeDescriptions.append(_item3)

    def _objectGIDFunc(self, input: CustomDataWrapper):
        self.objectGID = input.readVarUhInt()
        if self.objectGID < 0:
            raise ValueError("Forbidden value (" + str(self.objectGID) + ") on element of ExchangeTypesItemsExchangerDescriptionForUserMessage.objectGID.")
      
    def _objectTypeFunc(self, input: CustomDataWrapper):
        self.objectType = input.readInt()
        if self.objectType < 0:
            raise ValueError("Forbidden value (" + str(self.objectType) + ") on element of ExchangeTypesItemsExchangerDescriptionForUserMessage.objectType.")
    
    def __str__(self):
        if self.objectType in all_items:
            return f"Object: all_items={all_items[self.objectType]}, itemTypeDescriptions: [\n" + '\n'.join([str(i) for i in self.itemTypeDescriptions]) + "\n]"
        else:
            return f"Object: {self.objectType}, itemTypeDescriptions: [\n" + '\n'.join([str(i) for i in self.itemTypeDescriptions]) + "\n]"