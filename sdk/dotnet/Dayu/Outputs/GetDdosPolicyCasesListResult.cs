// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Dayu.Outputs
{

    [OutputType]
    public sealed class GetDdosPolicyCasesListResult
    {
        /// <summary>
        /// App protocol set of the DDoS policy case.
        /// </summary>
        public readonly ImmutableArray<string> AppProtocols;
        /// <summary>
        /// App type of the DDoS policy case.
        /// </summary>
        public readonly string AppType;
        /// <summary>
        /// Create time of the DDoS policy case.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Indicate whether the service involves overseas or not.
        /// </summary>
        public readonly string HasAbroad;
        /// <summary>
        /// Indicate whether the service actively initiates TCP requests or not.
        /// </summary>
        public readonly string HasInitiateTcp;
        /// <summary>
        /// Indicate whether the actively initiate UDP requests or not.
        /// </summary>
        public readonly string HasInitiateUdp;
        /// <summary>
        /// Indicate whether the service involves VPN service or not.
        /// </summary>
        public readonly string HasVpn;
        /// <summary>
        /// The max length of TCP message package.
        /// </summary>
        public readonly string MaxTcpPackageLen;
        /// <summary>
        /// The max length of UDP message package.
        /// </summary>
        public readonly string MaxUdpPackageLen;
        /// <summary>
        /// The minimum length of TCP message package.
        /// </summary>
        public readonly string MinTcpPackageLen;
        /// <summary>
        /// The minimum length of UDP message package.
        /// </summary>
        public readonly string MinUdpPackageLen;
        /// <summary>
        /// Name of the DDoS policy case.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The port that actively initiates TCP requests.
        /// </summary>
        public readonly string PeerTcpPort;
        /// <summary>
        /// The port that actively initiates UDP requests.
        /// </summary>
        public readonly string PeerUdpPort;
        /// <summary>
        /// Platform set of the DDoS policy case.
        /// </summary>
        public readonly ImmutableArray<string> PlatformTypes;
        /// <summary>
        /// Type of the resource that the DDoS policy case works for, valid values are `bgpip`, `bgp`, `bgp-multip` and `net`.
        /// </summary>
        public readonly string ResourceType;
        /// <summary>
        /// ID of the DDoS policy case to be query.
        /// </summary>
        public readonly string SceneId;
        /// <summary>
        /// End port of the TCP service.
        /// </summary>
        public readonly string TcpEndPort;
        /// <summary>
        /// The fixed signature of TCP protocol load.
        /// </summary>
        public readonly string TcpFootprint;
        /// <summary>
        /// Start port of the TCP service.
        /// </summary>
        public readonly string TcpStartPort;
        /// <summary>
        /// End port of the UDP service.
        /// </summary>
        public readonly string UdpEndPort;
        /// <summary>
        /// The fixed signature of TCP protocol load.
        /// </summary>
        public readonly string UdpFootprint;
        /// <summary>
        /// Start port of the UDP service.
        /// </summary>
        public readonly string UdpStartPort;
        /// <summary>
        /// Web API url set.
        /// </summary>
        public readonly ImmutableArray<string> WebApiUrls;

        [OutputConstructor]
        private GetDdosPolicyCasesListResult(
            ImmutableArray<string> appProtocols,

            string appType,

            string createTime,

            string hasAbroad,

            string hasInitiateTcp,

            string hasInitiateUdp,

            string hasVpn,

            string maxTcpPackageLen,

            string maxUdpPackageLen,

            string minTcpPackageLen,

            string minUdpPackageLen,

            string name,

            string peerTcpPort,

            string peerUdpPort,

            ImmutableArray<string> platformTypes,

            string resourceType,

            string sceneId,

            string tcpEndPort,

            string tcpFootprint,

            string tcpStartPort,

            string udpEndPort,

            string udpFootprint,

            string udpStartPort,

            ImmutableArray<string> webApiUrls)
        {
            AppProtocols = appProtocols;
            AppType = appType;
            CreateTime = createTime;
            HasAbroad = hasAbroad;
            HasInitiateTcp = hasInitiateTcp;
            HasInitiateUdp = hasInitiateUdp;
            HasVpn = hasVpn;
            MaxTcpPackageLen = maxTcpPackageLen;
            MaxUdpPackageLen = maxUdpPackageLen;
            MinTcpPackageLen = minTcpPackageLen;
            MinUdpPackageLen = minUdpPackageLen;
            Name = name;
            PeerTcpPort = peerTcpPort;
            PeerUdpPort = peerUdpPort;
            PlatformTypes = platformTypes;
            ResourceType = resourceType;
            SceneId = sceneId;
            TcpEndPort = tcpEndPort;
            TcpFootprint = tcpFootprint;
            TcpStartPort = tcpStartPort;
            UdpEndPort = udpEndPort;
            UdpFootprint = udpFootprint;
            UdpStartPort = udpStartPort;
            WebApiUrls = webApiUrls;
        }
    }
}
