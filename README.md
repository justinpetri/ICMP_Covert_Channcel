# ICMP_Covert_Channcel
This project sends an ICMP echo message from device to another. Within the sent packet is covert message (base64 encoded text). The second device runs the listener to capture the packet then decode the message. The main priorities for this project were to further my knowledge with golang and understanding of ICMP packets.


Separate files into two separate machines (or just use 127.0.0.1 as IP address for targetIP)
change targetIP
change stringToEncode


USE:
- change "ENTER_TARGET_IP_HERE" to an IP in the ICMP_sender.go file
- run ICMP_receiver.go to begin packet capture
- run ICMP_sender.go

SOURCES USED:
- https://pkg.go.dev/golang.org/x/net/icmp#PacketConn

NOTES:
- first 8 bits of received data from ICMP packet is reserved: 1 bit for type, 1 for code, 2 for checksum, 2 for identifier, 2 for sequence number
- ALTERNATIVE: storage channel modifying the data of the packet (Change last bit of code or checksum. split up encoded message and send 1 bit per packet (see if i change different fields like sequence number or type)
