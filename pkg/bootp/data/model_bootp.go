package data

import "net"

// MinSizeMessage this is minimal octet message
// See: RFC 1542 section 2.1
//  The IP Total Length and UDP Length must be large enough to
//  contain the minimal BOOTP header of 300 octets (in the UDP
//  data field) specified in [1].
//
// Link: https://tools.ietf.org/html/rfc1542#section-2.1
const MinSizeMessage uint = 300

// BOOTP this is data model for BOOTP protocol
// See: RFC 951 and RFC 1542 for details.
type BOOTP struct {
	Op     OperationCode    // Op - Code of operation
	HType  HardwareType     // HType  - Type of hardware address
	HLen   int8             // HLen   - Len of hardware address
	Hops   int8             // Hops   - Count of gateway between client and server
	Sec    int16            // Sec    - Number of seconds since initialization started
	Flags  int16            // Flags  - Flag of type message, broadcast or unicast (Not used)
	CIADDR net.IP           // CIADDR - The client's ip address, if the client knows it
	YIADDR net.IP           // YIADDR - Ip address for client
	SIADDR net.IP           // SIADDR - Ip address of BOOTP server
	GIADDR net.IP           // GIADDR - Ip address of relay gateway
	CHADDR net.HardwareAddr // CHADDR - Hardware address of client
	SName  string           // SName  - Name of server (Is optional field)
	File   string           // File   - Boot file name
	Vend   []Vendor         // Vend   - Is optional fields of vendor's data
}
