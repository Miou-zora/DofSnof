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
