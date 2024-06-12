package tracing

// Stable OTel Attributes
// See https://opentelemetry.io/docs/specs/semconv/general/attributes/
const (
	AttributeServerAddress          = "server.address"
	AttributeServerPort             = "server.port"
	AttributeClientAddress          = "client.address"
	AttributeClientPort             = "client.port"
	AttributeSourceAddress          = "source.address"
	AttributeSourcePort             = "source.port"
	AttributeDestinationAddress     = "destination.address"
	AttributeDestinationPort        = "destination.port"
	AttributeNetworkLocalAddress    = "network.local.address"
	AttributeNetworkLocalPort       = "network.local.port"
	AttributeNetworkPeerAddress     = "network.peer.address"
	AttributeNetworkPeerPort        = "network.peer.port"
	AttributeNetworkProtocolName    = "network.protocol.name"
	AttributeNetworkProtocolVersion = "network.protocol.version"
	AttributeNetworkTransport       = "network.transport"
	AttributeNetworkType            = "network.type"
	AttributeServiceName            = "service.name"
	AttributeServiceVersion         = "service.version"
	AttributeTelemetrySdkLanguage   = "telemetry.sdk.language"
	AttributeTelemetrySdkName       = "telemetry.sdk.name"
	AttributeTelemetrySdkVersion    = "telemetry.sdk.version"

	NetworkTypeTcp  = "tcp"
	NetworkTypeUdp  = "udp"
	NetworkTypePipe = "pipe"
	NetworkTypeUnix = "unix"
	NetworkTypeIpv4 = "ipv4"
	NetworkTypeIpv6 = "ipv6"
)

// Span OTel Attributes
const (
	AttributeSpanKind      = "span.kind"
	AttributeSpanNumEvents = "span.num_events"
	AttributeSpanNumLinks  = "span.num_links"
)

// HTTP Attributes
const (
	AttributeHttpHost       = "http.host"
	AttributeHttpMethod     = "http.method"
	AttributeHttpPath       = "http.path"
	AttributeHttpScheme     = "http.scheme"
	AttributeHttpStatusCode = "http.status_code"
	AttributeHttpUrl        = "http.url"
	AttributeHttpUserAgent  = "http.user_agent"
	AttributeHttpClientIp   = "http.client_ip"
)

// Common Attributes
const (
	AttributeErrorMessage = "error_msg"
	AttributeHasError     = "error" // should be 'true' or 'false'
	AttributeErrorType    = "error_type"
)
