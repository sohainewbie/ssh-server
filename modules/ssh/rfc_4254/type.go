package rfc_4254

type ChannelRequestType string

const (
	PtyReqType       ChannelRequestType = "pty-req"
	X11ReqType       ChannelRequestType = "x11-req"
	EnvType          ChannelRequestType = "env"
	ShellType        ChannelRequestType = "shell"
	ExecType         ChannelRequestType = "exec"
	SubsystemType    ChannelRequestType = "subsystem"
	WindowChangeType ChannelRequestType = "window-change"
	XonXoffType      ChannelRequestType = "xon-xoff"
	SignalType       ChannelRequestType = "signal"
	ExitStatusType   ChannelRequestType = "exit-status"
	ExitSignalType   ChannelRequestType = "exit-signal"
)

type ChannelOpenType string

const (
	ForwardedTCPIPType ChannelOpenType = "forwarded-tcpip"
	DirectTCPIPType    ChannelOpenType = "direct-tcpip"
	SessionType        ChannelOpenType = "session"
	X11Type            ChannelOpenType = "x11"
)

type GlobalRequestType string

const (
	TcpipForwardType       GlobalRequestType = "tcpip-forward"
	CancelTcpipForwardType GlobalRequestType = "cancel-tcpip-forward"
)

/*
 "" means that connections are to be accepted on all protocol families supported by the SSH implementation.

 "0.0.0.0" means to listen on all IPv4 addresses.

 "::" means to listen on all IPv6 addresses.

 "localhost" means to listen on all protocol families supported by the SSH implementation on loopback addresses only ([RFC3330] and [RFC3513]).

 "127.0.0.1" and "::1" indicate listening on the loopback interfaces for IPv4 and IPv6, respectively.

*/
