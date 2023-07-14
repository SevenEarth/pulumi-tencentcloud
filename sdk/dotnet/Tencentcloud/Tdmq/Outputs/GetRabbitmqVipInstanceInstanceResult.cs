// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tdmq.Outputs
{

    [OutputType]
    public sealed class GetRabbitmqVipInstanceInstanceResult
    {
        public readonly int AutoRenewFlag;
        public readonly string ConfigDisplay;
        public readonly string ExceptionInformation;
        public readonly int ExpireTime;
        public readonly string InstanceId;
        public readonly string InstanceName;
        public readonly string InstanceVersion;
        public readonly int MaxBandWidth;
        public readonly int MaxStorage;
        public readonly int MaxTps;
        public readonly int NodeCount;
        public readonly int PayMode;
        public readonly string Remark;
        public readonly string SpecName;
        public readonly int Status;

        [OutputConstructor]
        private GetRabbitmqVipInstanceInstanceResult(
            int autoRenewFlag,

            string configDisplay,

            string exceptionInformation,

            int expireTime,

            string instanceId,

            string instanceName,

            string instanceVersion,

            int maxBandWidth,

            int maxStorage,

            int maxTps,

            int nodeCount,

            int payMode,

            string remark,

            string specName,

            int status)
        {
            AutoRenewFlag = autoRenewFlag;
            ConfigDisplay = configDisplay;
            ExceptionInformation = exceptionInformation;
            ExpireTime = expireTime;
            InstanceId = instanceId;
            InstanceName = instanceName;
            InstanceVersion = instanceVersion;
            MaxBandWidth = maxBandWidth;
            MaxStorage = maxStorage;
            MaxTps = maxTps;
            NodeCount = nodeCount;
            PayMode = payMode;
            Remark = remark;
            SpecName = specName;
            Status = status;
        }
    }
}
