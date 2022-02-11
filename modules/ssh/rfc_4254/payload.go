package rfc_4254

import (
	"golang.org/x/crypto/ssh"
)

// See RFC 4254, section 7.2
type ForwardedTCPIP struct {
	//senderChannel uint32
	//initialWindowSize uint32
	//maximumPacketSize uint32
	Addr       string
	Port       uint32
	OriginAddr string
	OriginPort uint32
}

// See RFC 4254, section 7.2
type DirectTCP struct {
	//senderChannel uint32
	//initialWindowSize uint32
	//maximumPacketSize uint32
	Host       string
	Port       uint32
	OriginAddr string
	OriginPort uint32
}

type Session struct{}

// See RFC 4254, section 6.3.2
type X11 struct {
	//senderChannel uint32
	//initialWindowSize uint32
	//maximumPacketSize uint32
	OriginAddr string
	OriginPort uint32
}

// See RFC 4254, section 6.2
type PtyReq struct {
	//recipientChannel uint32
	//type string
	//wantReply boolean
	TermEnvVar     string //TERM environment variable value (e.g., vt100)
	WidthInChars   uint32 //terminal width, characters (e.g., 80)
	HeightInRows   uint32 //terminal height, rows (e.g., 24)
	WidthInPixels  uint32 //terminal width, pixels (e.g., 640)
	HeightInPixels uint32 //terminal height, pixels (e.g., 480)
	TermModes      string //encoded terminal modes
}

// See RFC 4254, section 6.3.1
type X11Req struct {
	//recipientChannel uint32
	//type string
	//wantReply boolean
	SingleConnection bool   //single connection
	AuthProtocol     string // x11 authentication protocol
	AuthCookie       string // x11 authentication cookie
	ScreenNr         uint32 // x11 screen number
}

// See RFC 4254, section 6.4
type Env struct {
	//recipientChannel uint32
	//type string
	//wantReply boolean
	Name  string //variable name
	Value string //variable value
}

type Shell struct{}
type Exec struct{ Command string }
type Subsystem struct{ Name string }

type WindowChange struct {
	//byte      SSH_MSG_CHANNEL_REQUEST
	//uint32    recipient channel
	//string    "window-change"
	//boolean   FALSE
	WidthColumns uint32 //terminal width, columns
	HeightRows   uint32 //terminal height, rows
	WidthPixels  uint32 //terminal width, pixels
	HeightPixels uint32 //terminal height, pixels
}

// See RFC 4254, section 6.8
type XonXoff struct {
	ClientCanDo bool
}

// See RFC 4254, section 6.9
type Signal struct {
	//byte      SSH_MSG_CHANNEL_REQUEST
	//uint32    recipient channel
	//string    "window-change"
	//boolean   FALSE
	SignalName ssh.Signal //signal name (without the "SIG" prefix)
}

type ExitStatus struct {
	ExitStatus uint32
}

type ExitSignal struct {
	SignalName ssh.Signal //signal name (without the "SIG" prefix)
	CoreDumped bool       //core dumped
	Error      string     //error message in ISO-10646 UTF-8 encoding
	Language   string     //language tag [RFC3066]}
}

// See RFC 4254, section 7.1
type TcpipForward struct {
	//byte      SSH_MSG_GLOBAL_REQUEST
	//string    "tcpip-forward"
	//boolean   want reply
	Addr string //address_to_bind (e.g., "127.0.0.1")
	Port uint32 //port number to bind
}

// See RFC 4254, section 7.1
type CancelTcpipForward struct {
	//byte      SSH_MSG_GLOBAL_REQUEST
	//string    "tcpip-forward"
	//boolean   want reply
	Addr string //address_to_bind (e.g., "127.0.0.1")
	Port uint32 //port number to bind
}

func ParseRequest(request *ssh.Request) (interface{}, error) {
	switch request.Type {
	case "pty-req":
		var payload PtyReq
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "x11-req":
		var payload X11Req
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "env":
		var payload Env
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "shell":
		var payload Shell
		return payload, nil
	case "exec":
		var payload Exec
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "subsystem":
		var payload Subsystem
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "window-change":
		var payload WindowChange
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "xon-xoff":
		var payload XonXoff
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "signal":
		var payload Signal
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "exit-status":
		var payload ExitStatus
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "exit-signal":
		var payload ExitSignal
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "tcpip-forward":
		var payload TcpipForward
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err
	case "cancel-tcpip-forward":
		var payload CancelTcpipForward
		err := ssh.Unmarshal(request.Payload, &payload)
		return payload, err

	default:
		return nil, nil
	}
}

func ParseNewChannel(nchn ssh.NewChannel) (interface{}, error) {
	switch nchn.ChannelType() {

	case "forwarded-tcpip":
		var payload ForwardedTCPIP
		err := ssh.Unmarshal(nchn.ExtraData(), &payload)
		return payload, err
	case "direct-tcpip":
		var payload DirectTCP
		err := ssh.Unmarshal(nchn.ExtraData(), &payload)
		return payload, err
	case "session":
		var payload Session
		return payload, nil
	case "x11":
		var payload X11
		err := ssh.Unmarshal(nchn.ExtraData(), &payload)
		return payload, err

	default:
		return nil, nil
	}
}
