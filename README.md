# ICMP_Covert_Channcel
This project sends an ICMP echo message from device to another. Within the sent packet is covert message (base64 encoded text). The second device runs the listener to capture the packet then decode the message. The main priorities for this project were to further my knowledge with golang and understanding of ICMP packets.

USE:
- change

SOURCES USED:
- https://blog.logrocket.com/learn-golang-encryption-decryption/
- https://stackoverflow.com/questions/2937123/implementing-icmp-ping-in-go
- https://pkg.go.dev/golang.org/x/net/icmp#PacketConn

NOTES:
first 8 bits of received data from ICMP packet is reserved: 1 bit for type, 1 for code, 2 for checksum, 2 for identifier, 2 for sequence number
