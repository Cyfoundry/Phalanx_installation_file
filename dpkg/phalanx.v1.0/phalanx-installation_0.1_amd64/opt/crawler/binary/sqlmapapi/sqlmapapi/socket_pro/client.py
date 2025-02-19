import socket
import os

class Client():
    def __init__(self, host, port, buff_size) -> None:
        self.host = host
        self.port = port
        self.buff_size = buff_size
    def transfer(self, payload):
        # create the client socket
        s = socket.socket()
        #print(f"[+] Connecting to {self.host}:{self.port}")
        s.connect((self.host, self.port))
        #print("[+] Connected.")
        # send the files to the server
        if isinstance(payload, str): 
            payload = payload.encode()
            s.sendall(payload)
        file_name = s.recv(1024)
        s.close()
        return file_name
    def remove_file(self, filename):
        s = socket.socket()
        s.connect((self.host, self.port))
        s.send(f"{filename}".encode())
        s.close()


if __name__ == "__main__":
    client = Client('192.168.10.219', 9898, 4096)
    file_name = client.transfer('POST lib/login.php HTTP/1.1\nHost: 192.168.10.188')
    print(os.path.basename(file_name))
