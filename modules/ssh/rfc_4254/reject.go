package rfc_4254

import "golang.org/x/crypto/ssh"

const (
	SshOpenAdministrativelyProhibited ssh.RejectionReason = 1
	SshOpenConnectFailed              ssh.RejectionReason = 2
	SshOpenUnknownChannelType         ssh.RejectionReason = 3
	SshOpenResourceShortage           ssh.RejectionReason = 4
)
