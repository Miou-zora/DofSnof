###
### This code is useless, it's just a inheritance of the original code
### It also contains some dark code (Sniffer.receive) that I don't understand
###

from scapy.all import sniff, Raw, IP, ICMP # pylint: disable=no-name-in-module
from colorama import Fore, Back, Style
from CustomDataWrapper import Data, Buffer
from ProtocolBuilder import ProtocolBuilder
from Misc import * # pylint: disable=unused-wildcard-import
import struct
from all_message import all_message
from all_type import all_types

all_items = {}

class CustomDataWrapper:
    def __init__(self, data = None):
        self.data = data if data else bytearray()

    def writeByte(self, value):
        self.data += value.to_bytes(1, byteorder='big')

    def writeShort(self, value):
        self.data += value.to_bytes(2, byteorder='big')

    def writeInt(self, value):
        self.data += value.to_bytes(4, byteorder='big')

    def writeDouble(self, value):
        self.data += value.to_bytes(8, byteorder='big')

    def writeUTF(self, value):
        self.writeShort(len(value))
        self.data += value.encode('utf-8')

    def readByte(self):
        return int.from_bytes(self.data[:1], byteorder='big')

    def readShort(self):
        return int.from_bytes(self.data[:2], byteorder='big')
    
    def readUnsignedShort(self):
        return int.from_bytes(self.data[:2], byteorder='big')

    def readInt(self):
        return int.from_bytes(self.data[:4], byteorder='big')

    def readDouble(self):
        return int.from_bytes(self.data[:8], byteorder='big')
    
    def readVarUhLong(self):
        value = 0
        i = 0
        while i < 64:
            byte = self.readByte()
            value = value | (byte & 127) << i
            if byte & 128 == 0:
                return value
            i += 7
        raise ValueError("Too much data")
    
    def readVarUhInt(self):
        value = 0
        i = 0
        while i < 32:
            byte = self.readByte()
            value = value | (byte & 127) << i
            if byte & 128 == 0:
                return value
            i += 7
        raise ValueError("Too much data")
    
    def readVarUhShort(self):
        value = 0
        i = 0
        while i < 16:
            byte = self.readByte()
            value = value | (byte & 127) << i
            if byte & 128 == 0:
                return value
            i += 7
        raise ValueError("Too much data")

    def readUTF(self):
        length = self.readShort()
        value = self.data[2:2+length].decode('utf-8')
        self.data = self.data[2+length:]
        return value

class ChatAbstractServerMessage:
    def __init__(self, channel = 0, content = '', timestamp = 0, fingerprint = ''):
        self.channel = channel
        self.content = content
        self.timestamp = timestamp
        self.fingerprint = fingerprint

    def serialize(self, output: CustomDataWrapper):
        output.writeByte(self.channel)
        output.writeUTF(self.content)
        if self.timestamp < 0:
            raise ValueError("Forbidden value (" + str(self.timestamp) + ") on element timestamp.")
        output.writeInt(self.timestamp)
        output.writeUTF(self.fingerprint)

    def deserialize(self, input: Buffer):
        self.channel = input.readByte()
        self.content = input.readUTF()
        self.timestamp = input.readInt()
        self.fingerprint = input.readUTF()

    def getMessageId(self) -> int:
        return 880

class ChatServerMessage(ChatAbstractServerMessage):
    def __init__(self, channel = 0, content = '', timestamp = 0, fingerprint = '', senderId = 0, senderName = '', prefix = '', senderAccountId = 0):
        super().__init__(channel, content, timestamp, fingerprint)
        self.senderId = senderId
        self.senderName = senderName
        self.prefix = prefix
        self.senderAccountId = senderAccountId

    def serialize(self, output: CustomDataWrapper):
        super().serialize(output)
        if self.senderId < -9007199254740990 or self.senderId > 9007199254740990:
            raise ValueError("Forbidden value (" + str(self.senderId) + ") on element senderId.")
        output.writeDouble(self.senderId)
        output.writeUTF(self.senderName)
        output.writeUTF(self.prefix)
        if self.senderAccountId < 0:
            raise ValueError("Forbidden value (" + str(self.senderAccountId) + ") on element senderAccountId.")
        output.writeInt(self.senderAccountId)

    def deserialize(self, input):
        super().deserialize(input)
        self.senderId = input.readDouble()
        self.senderName = input.readUTF()
        self.prefix = input.readUTF()
        self.senderAccountId = input.readInt()

    def getMessageId(self) -> int:
        return 881
    
    def __str__(self) -> str:
        return f"channel: {self.channel}, content: {self.content}, timestamp: {self.timestamp}, fingerprint: {self.fingerprint}, senderId: {self.senderId}, senderName: {self.senderName}, prefix: {self.prefix}, senderAccountId: {self.senderAccountId}"

class ObjectEffect:
    def __init__(self, actionId = 0) -> None:
        self.actionId = actionId

    def deserialize(self, input: CustomDataWrapper):
        self.actionId = input.readVarUhShort()

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
        # return f"objectGID: {self.objectGID}, objectType: {self.objectType}, itemTypeDescriptions: [" + ', '.join([str(i) for i in self.itemTypeDescriptions]) + "]"
        if self.objectType in all_items:
            return f"Object: all_items={all_items[self.objectType]}, itemTypeDescriptions: [\n" + '\n'.join([str(i) for i in self.itemTypeDescriptions]) + "]"
        else:
            return f"Object: {self.objectType}, itemTypeDescriptions: [\n" + '\n'.join([str(i) for i in self.itemTypeDescriptions]) + "]"

class BasicPongMessage:
    # store nothing and deserialize is only a boolean
    def __init__(self) -> None:
        self.quiet = True

    def deserialize(self, input: CustomDataWrapper):
        self.quiet = input.readByte() == 1
    
    def __str__(self):
        return f"BasicPongMessage: quiet: {self.quiet}"
    
    def getMessageId(self) -> int:
        return 4877


id_to_class = {
    1770: ChatAbstractServerMessage,
    6772: ChatServerMessage,
    2738: ExchangeTypesItemsExchangerDescriptionForUserMessage,
    6572: ExchangeTypesExchangerDescriptionForUserMessage,
    4877: BasicPongMessage
}

class Msg():
    def __init__(self, buffer: Buffer, protocol):
        self.b = True
        self.protocol = protocol
        self.error = ''
        print("Buffer: ", buffer)
        try:
            header = int.from_bytes(buffer.read(2), byteorder="big")
            self.id = header >> 2
            self.lenType = header & 3
            self.dataLen = int.from_bytes(buffer.read(self.lenType), byteorder="big")
            self.checkHeader()
            deserialized = False
            if (self.id in all_message):
                if (self.id in id_to_class):
                    value = id_to_class[self.id]()
                    try:
                        aze = Buffer(buffer.read(self.dataLen))
                        buffer.end()
                        value.deserialize(aze)
                    except Exception as e:
                        # print stack trace
                        eprint(f"Exception: {e}")
                    deserialized = True
                    sprint(f"\t{value}")
                a=self.id
            else:
                eprint(f"Message not recognised id={self.id} from {len(buffer.data)} bytes")
            try:
                if (buffer.pos == len(buffer.data)):
                    pass
                elif not deserialized:
                    self.data = Data(buffer.read(self.dataLen))
            except Exception as e:
                eprint(f"Exception1: {e}")
            deserialized = False
        except IndexError as e:
            eprint(f"IndexError: {e}")
            buffer.pos = 0
            self.b = False
        except ValueError:
            eprint(f"ValueError: {self.error}")
            buffer.pos = 0
            self.b = False
        except Exception as e:
            eprint(f"Exception2: {e}")
        else:
            buffer.end()

    def checkHeader(self):
        if not self.lenType in [0, 1, 2, 3]:
            self.error = 'Wrong lenType "' + str(self.lenType) + '"'
            raise ValueError

    def __bool__(self):
        return self.b

class Sniffer:
    def __init__(self, concatMode = True):
        self.protocolBuilder = ProtocolBuilder()
        self.protocol = self.protocolBuilder.protocol
        self.buffer = Buffer()
        self.concatMode = concatMode
        self.lastPkt = None

    def run(self, callback, whitelist = None):
        self.callback = callback
        self.whitelist = whitelist
        sniff(
            filter='tcp src port 5555',
            lfilter = lambda pkt: pkt.haslayer(Raw),
            prn = lambda pkt: self.receive(pkt)
        )

    def receive(self, pkt):
        # I don't know what is this
        ### From This
        if self.lastPkt and pkt.getlayer(IP).src != self.lastPkt.getlayer(IP).src:
            self.lastPkt = None
        if self.lastPkt and pkt.getlayer(IP).id < self.lastPkt.getlayer(IP).id:
            self.buffer.reorder(bytes(pkt.getlayer(Raw)),
            len(self.lastPkt.getlayer(Raw)))
        else:
            if self.concatMode:
                self.buffer += bytes(pkt.getlayer(Raw))
            else:
                self.buffer = Buffer()
                self.buffer += bytes(pkt.getlayer(Raw))
        self.lastPkt = pkt
        ### To This


        while len(self.buffer) and Msg(self.buffer, self.protocol):
            pass
        