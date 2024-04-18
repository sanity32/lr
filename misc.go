package lr

import "errors"

const DEFAULT_RPC_SERVER_PORT int = 22217

var ErrClientNotInit = errors.New("client is not initialized")
