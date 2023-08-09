# ICMP_Covert_Channcel
The goal of this project was to develop an understanding of Golang and Covert (Timing) Channels. A Covert Timing Channel does not modify the packets but instead sends message at specific time intervals, allowing those discrepencies to be used to form a message. In this case, the user can enter a string and the program will convert it into binary. ICMP packets will then be sent at specific intervals to match the encoded binary message. The receiver will be listening and will pick up the ICMP packets, identify the discrepencies in time, convert to binary and then decode the hidden message. The timing intervals can be increased in order to slow down the sending rate of ICMP packet. Additionally, this tool can be repurposed to take in messages from other files.

This project is for educational purposes only!


USE:
- Separate ICMP_receiver.go and ICMP_sender.go onto two different machines (or run locally by entering 127.0.0.1 for targetIP)
- change targetIP value in ICMP_sender.go
- change stringToEncode value in ICMP_sender.go
- Run ICMP_receiver.go before running ICMP_sender.go


SOURCES USED:
- https://pkg.go.dev/golang.org/x/net/icmp#PacketConn


NOTES:
- first 8 bits of received data from ICMP packet is reserved: 1 bit for type, 1 for code, 2 for checksum, 2 for identifier, 2 for sequence number
- ALTERNATIVE: storage channel modifying the data of the packet (Change last bit of code or checksum. split up encoded message and send 1 bit per packet (see if i change different fields like sequence number or type)
