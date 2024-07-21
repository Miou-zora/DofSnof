from Buffer import Buffer
from typing import Callable
from scapy.all import AsyncSniffer, Packet, Raw
import threading

class PacketReceiver:
    def __init__(self, filter: str = 'tcp src port 5555', lfilter: Callable[[any], None] = lambda pkt: pkt.haslayer(Raw), max_buffer_size = 4096) -> None:
        self.buffer: Buffer = Buffer()
        self._filter = filter
        self._lfilter = lfilter
        self._thread: threading.Thread = None
        self._max_buffer_size: int = max_buffer_size

    def run(self) -> None:
        self._thread = AsyncSniffer(
            filter=self._filter,
            lfilter=self._lfilter,
            prn=lambda pkt: self.__receive(pkt)
        )
        self._thread.start()

    def stop(self) -> None:
        if self._thread is not None:
            self._thread.stop()
            self._thread = None
        
    def __receive(self, pkt: Packet):
        new_packet: bytes = pkt[Raw].load
        if len(self.buffer) + len(new_packet) > self._max_buffer_size:
            print(f"Buffer is full, resetting buffer")
            self.buffer = []
            return
        new_packet = bytearray(new_packet)
        self.buffer += new_packet