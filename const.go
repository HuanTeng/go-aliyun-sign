package sign

// HTTP Header keys
const (
	HTTPHeaderAccept      = "Accept"
	HTTPHeaderContentMD5  = "Content-MD5"
	HTTPHeaderContentType = "Content-Type"
	HTTPHeaderUserAgent   = "User-Agent"
	HTTPHeaderDate        = "Date"
)

// HTTP Header keys used for aliyun signature
const (
	HTTPHeaderCAPrefix           = "X-Ca-"
	HTTPHeaderCASignature        = "X-Ca-Signature"
	HTTPHeaderCATimestamp        = "X-Ca-Timestamp"
	HTTPHeaderCANonce            = "X-Ca-Nonce"
	HTTPHeaderCAKey              = "X-Ca-Key"
	HTTPHeaderCASignatureHeaders = "X-Ca-Signature-Headers"
)

const (
	defaultUserAgent = "Go-Aliyun-Client"
)
