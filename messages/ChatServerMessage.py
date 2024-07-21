from .ChatAbstractServerMessage import ChatAbstractServerMessage
from CustomDataWrapper import CustomDataWrapper

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
    