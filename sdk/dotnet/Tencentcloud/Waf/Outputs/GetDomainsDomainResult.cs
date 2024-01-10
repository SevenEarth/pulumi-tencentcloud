// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainResult
    {
        /// <summary>
        /// Traffic Source: clb represents Tencent Cloud clb, apisix represents apisix gateway, tsegw represents Tencent Cloud API gateway, default clbNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string AlbType;
        /// <summary>
        /// API security switch status, 0 off, 1 onNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int ApiStatus;
        /// <summary>
        /// User appid.
        /// </summary>
        public readonly int AppId;
        /// <summary>
        /// BOT switch status, 0 off, 1 on.
        /// </summary>
        public readonly int BotStatus;
        /// <summary>
        /// Waf sandbox export addresses, should be added to the whitelist by the upstreams.
        /// </summary>
        public readonly ImmutableArray<string> CcLists;
        /// <summary>
        /// Cdc clustersNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string CdcClusters;
        /// <summary>
        /// Whether to enable access logs, 1 enable, 0 disable.
        /// </summary>
        public readonly int ClsStatus;
        /// <summary>
        /// Cname address, used for dns access.
        /// </summary>
        public readonly string Cname;
        /// <summary>
        /// Create time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Domain name.
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// Domain unique ID.
        /// </summary>
        public readonly string DomainId;
        /// <summary>
        /// Instance type, sparta-waf represents SAAS WAF, clb-waf represents CLB WAF.
        /// </summary>
        public readonly string Edition;
        /// <summary>
        /// Rule and AI Defense Mode, 10 Rule Engine Observation&amp;amp;amp;&amp;amp;amp;AI Engine Shutdown Mode 11 Rule Engine Observation&amp;amp;amp;&amp;amp;amp;AI Engine Observation Mode 12 Rule Engine Observation&amp;amp;amp;&amp;amp;amp;AI Engine Interception Mode 20 Rule Engine Interception&amp;amp;amp;&amp;amp;amp;AI Engine Shutdown Mode 21 Rule Engine Interception&amp;amp;amp;&amp;amp;amp;AI Engine Observation Mode 22 Rule Engine Interception&amp;amp;amp;&amp;amp;amp;AI Engine Interception Mode.
        /// </summary>
        public readonly int Engine;
        /// <summary>
        /// CLBWAF traffic mode, 1 cleaning mode, 0 mirroring mode.
        /// </summary>
        public readonly int FlowMode;
        /// <summary>
        /// Unique ID of Instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Instance name.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// Ipv6 switch status, 0 off, 1 on.
        /// </summary>
        public readonly int Ipv6Status;
        /// <summary>
        /// Instance level.
        /// </summary>
        public readonly int Level;
        /// <summary>
        /// List of bound LB.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainLoadBalancerSetResult> LoadBalancerSets;
        /// <summary>
        /// Rule defense mode, 0 observation mode, 1 interception mode.
        /// </summary>
        public readonly int Mode;
        /// <summary>
        /// Listening ports.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainPortResult> Ports;
        /// <summary>
        /// Whether to enable the delivery of CKafka function, 0 off, 1 on.
        /// </summary>
        public readonly int PostCkafkaStatus;
        /// <summary>
        /// Whether to enable the delivery CLS function, 0 off, 1 on.
        /// </summary>
        public readonly int PostClsStatus;
        /// <summary>
        /// Region.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Waf engine export addresses, should be added to the whitelist by the upstreams.
        /// </summary>
        public readonly ImmutableArray<string> RsLists;
        /// <summary>
        /// Detailed explanation of security group statusNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly string SgDetail;
        /// <summary>
        /// Security group status, 0 does not display, 1 non Tencent cloud source site, 2 security group binding failed, 3 security group changedNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        public readonly int SgState;
        /// <summary>
        /// Clbwaf domain name listener status, 0 operation successful, 4 binding LB, 6 unbinding LB, 7 unbinding LB failed, 8 binding LB failed, 10 internal error.
        /// </summary>
        public readonly int State;
        /// <summary>
        /// Waf switch,0 off 1 on.
        /// </summary>
        public readonly int Status;

        [OutputConstructor]
        private GetDomainsDomainResult(
            string albType,

            int apiStatus,

            int appId,

            int botStatus,

            ImmutableArray<string> ccLists,

            string cdcClusters,

            int clsStatus,

            string cname,

            string createTime,

            string domain,

            string domainId,

            string edition,

            int engine,

            int flowMode,

            string instanceId,

            string instanceName,

            int ipv6Status,

            int level,

            ImmutableArray<Outputs.GetDomainsDomainLoadBalancerSetResult> loadBalancerSets,

            int mode,

            ImmutableArray<Outputs.GetDomainsDomainPortResult> ports,

            int postCkafkaStatus,

            int postClsStatus,

            string region,

            ImmutableArray<string> rsLists,

            string sgDetail,

            int sgState,

            int state,

            int status)
        {
            AlbType = albType;
            ApiStatus = apiStatus;
            AppId = appId;
            BotStatus = botStatus;
            CcLists = ccLists;
            CdcClusters = cdcClusters;
            ClsStatus = clsStatus;
            Cname = cname;
            CreateTime = createTime;
            Domain = domain;
            DomainId = domainId;
            Edition = edition;
            Engine = engine;
            FlowMode = flowMode;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Ipv6Status = ipv6Status;
            Level = level;
            LoadBalancerSets = loadBalancerSets;
            Mode = mode;
            Ports = ports;
            PostCkafkaStatus = postCkafkaStatus;
            PostClsStatus = postClsStatus;
            Region = region;
            RsLists = rsLists;
            SgDetail = sgDetail;
            SgState = sgState;
            State = state;
            Status = status;
        }
    }
}
