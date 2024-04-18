package lr

import _ "embed"

//go:embed "stable/lrsrv-windows-x64.exe.zip"
var LRSRV_STABLE_ZIPPED EmbedZipFile

const DEFAULT_LRSRV_FILENAME = "lrsrv-windows-x64.exe"
