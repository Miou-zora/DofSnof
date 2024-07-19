from CustomDataWrapper import CustomDataWrapper

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