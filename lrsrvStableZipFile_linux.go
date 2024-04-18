package lr

import _ "embed"

//go:embed "stable/lrsrv-debian-x64.zip"
var LRSRV_STABLE_ZIPPED EmbedZipFile

const DEFAULT_LRSRV_FILENAME = "lrsrv-debian-x64"
