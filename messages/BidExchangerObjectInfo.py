from CustomDataWrapper import CustomDataWrapper
from .ObjectEffectInteger import ObjectEffectInteger
from Prints import eprint

class BidExchangerObjectInfo:
    def __init__(self) -> None:
        self.objectUID = 0
        self.objectGID = 0
        self.objectType = 0
        self.effects = []
        self.prices = []
      
    def deserialize(self, input: CustomDataWrapper):
        _id4 = 0
        _item4 = None
        _val5 = 0
        self._objectUIDFunc(input)
        self._objectGIDFunc(input)
        self._objectTypeFunc(input)
        _effectsLen = input.readUnsignedShort()

        for _ in range(_effectsLen):
            _id4 = input.readUnsignedShort()
            if _id4 == 3930:
                _item4 = ObjectEffectInteger()
                _item4.deserialize(input)
                self.effects.append(_item4)
            else:
                eprint("NOO ID4: ", _id4)

        _pricesLen = input.readUnsignedShort()
        for _ in range(_pricesLen):
            _val5 = input.readVarUhLong() # or readDouble
            if _val5 < 0 or _val5 > 9007199254740992:
                raise ValueError("Forbidden value (" + str(_val5) + ") on elements of prices.")
            self.prices.append(_val5)

    def _objectUIDFunc(self, input: CustomDataWrapper):
        self.objectUID = input.readVarUhInt()
        if self.objectUID < 0:
            raise ValueError("Forbidden value (" + str(self.objectUID) + ") on element of BidExchangerObjectInfo.objectUID.")
    
    def _objectGIDFunc(self, input: CustomDataWrapper):
        self.objectGID = input.readVarUhInt()
        if self.objectGID < 0:
            raise ValueError("Forbidden value (" + str(self.objectGID) + ") on element of BidExchangerObjectInfo.objectGID.")

    def _objectTypeFunc(self, input: CustomDataWrapper):
        self.objectType = input.readInt()
        if self.objectType < 0:
            raise ValueError("Forbidden value (" + str(self.objectType) + ") on element of BidExchangerObjectInfo.objectType.")
        
    def __str__(self):
        return f"objectUID: {self.objectUID}, objectGID: {self.objectGID}, objectType: {self.objectType}, effects: {[str(i) for i in self.effects]}, prices: {self.prices}"
