package packet

const (
	metalLBNamespace          = "metallb-system"
	metalLBConfigMapName      = "config"
	configMapResource         = "configmaps"
	hostnameKey               = "kubernetes.io/hostname"
	packetIdentifier          = "packet-ccm-auto"
	packetTag                 = "usage=" + packetIdentifier
	ccmIPDescription          = "Packet Kubernetes CCM auto-generated for Load Balancer"
	DefaultAnnotationNodeASN  = "packet.com/node-asn"
	DefaultAnnotationPeerASNs = "packet.com/peer-asn"
	DefaultAnnotationPeerIPs  = "packet.com/peer-ip"
	DefaultAnnotationSrcIP    = "packet.com/src-ip"
	DefaultLocalASN           = 65000
	DefaultPeerASN            = 65530
	DefaultAPIServerPort      = 6443
)
