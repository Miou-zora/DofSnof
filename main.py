from __future__ import annotations
from Buffer import Buffer
from all_message import all_message
from network.PacketReceiver import PacketReceiver
from messages.id_to_class import id_to_class
from Header import Header
from Prints import eprint, wprint, sprint

PR = PacketReceiver()
PR.run()
while True:
    if len(PR.buffer) >= 3:
        header: Header = Header.deserialize(PR.buffer.data)
        if not header.is_valid():
            eprint(f"Unknown packet id {header.id}")
            # what to do here ? Reset all packet ? Raise an error ?
        message = all_message[header.id]
        # print(f"Message received: {message}")
        size: int = header.data_len + header.len_type + 2
        if size > len(PR.buffer):
            wprint(f"Packet is not complete, waiting for more data...")
            continue
        packet = Buffer(PR.buffer.data[(2 + header.len_type):size])
        PR.buffer.data = PR.buffer.data[size:]
        print(f"Data: {packet}")
        print(f"Header: {header} - Message: {message}")
        if header.id in id_to_class:
            value = id_to_class[header.id]()
            value.deserialize(packet)
            sprint(f"Result: {value}")