from __future__ import annotations
from all_message import all_message

class Header:
    def __init__(self) -> None:
        self.id: int = -1
        self.len_type: int = -1
        self.data_len: int = -1

    @staticmethod
    def deserialize(input: bytearray) -> Header:
        if len(input) < 3:
            return None
        obj = Header()
        first_part = int.from_bytes(input[:2], "big")
        obj.id = first_part >> 2
        obj.len_type = first_part & 3
        obj.data_len = int.from_bytes(input[2:2+obj.len_type], "big")
        return obj
    
    def is_valid(self) -> bool:
        return self.id in all_message
    
    def __str__(self) -> str:
        return f"Header(id={self.id}, len_type={self.len_type}, data_len={self.data_len})"