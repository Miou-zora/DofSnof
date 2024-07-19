from CustomDataWrapper import CustomDataWrapper
from Buffer import Buffer

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
