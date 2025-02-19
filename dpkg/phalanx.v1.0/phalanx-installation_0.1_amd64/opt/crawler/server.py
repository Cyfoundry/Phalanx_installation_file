import socket
import tempfile
import os

# device's IP address
SERVER_HOST = "127.0.0.1"
SERVER_PORT = 9898
# receive 4096 bytes each time
BUFFER_SIZE = 4096
SEPARATOR = "<SEPARATOR>"

# create the server socket
# TCP socket
s = socket.socket()
# bind the socket to our local address
s.bind((SERVER_HOST, SERVER_PORT))
# enabling our server to accept connections
# 5 here is the number of unaccepted connections that
# the system will allow before refusing new connections
s.listen(100)
#print(f"[*] Listening as {SERVER_HOST}:{SERVER_PORT}")
count = 0
workdir = os.path.dirname(__file__)
# accept connection if there is any
while True:
    client_socket, address = s.accept()
    # if below code is executed, that means the sender is connected
    print(f"[+] {address} is connected.")

    # receive the string from the client
    string_data = client_socket.recv(BUFFER_SIZE).decode()
    ##remove the file if sqlmap scan finished.
    if 'txt' in string_data:
        os.remove(os.path.join(f'{workdir}/sqlmap/',string_data))
        print("Remove ", string_data, " successfully.")
    ## receive the file content of post methods
    else:
        # write the string to a temporary file
        f = open(os.path.join(f'{workdir}/sqlmap/',str(count)+'.txt'), 'w+')
        f.write(string_data)
        # send the file name back to client
        client_socket.send(f.name.encode())
        print("Finished writing ", f.name.encode())
        f.close()
        count += 1
    client_socket.close()

    # ## change the directory after creating temp file
# close the server socket
s.close()
