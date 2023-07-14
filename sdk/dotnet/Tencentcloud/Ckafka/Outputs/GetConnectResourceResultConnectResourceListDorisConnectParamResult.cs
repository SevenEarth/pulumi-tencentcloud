// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetConnectResourceResultConnectResourceListDorisConnectParamResult
    {
        /// <summary>
        /// Doris's http load balancing connection port, usually mapped to be's 8040 port.
        /// </summary>
        public readonly int BePort;
        /// <summary>
        /// Whether to update to the associated Dip task.
        /// </summary>
        public readonly bool IsUpdate;
        /// <summary>
        /// The password of the connection source.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// SQLServer port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Instance resource of connection source.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// Whether the connection source is a self-built cluster.
        /// </summary>
        public readonly bool SelfBuilt;
        /// <summary>
        /// Instance VIP of the connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        public readonly string ServiceVip;
        /// <summary>
        /// The vpc Id of the connection source, when it is a Tencent Cloud instance, it is required.
        /// </summary>
        public readonly string UniqVpcId;
        /// <summary>
        /// The username of the connection source.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetConnectResourceResultConnectResourceListDorisConnectParamResult(
            int bePort,

            bool isUpdate,

            string password,

            int port,

            string resource,

            bool selfBuilt,

            string serviceVip,

            string uniqVpcId,

            string userName)
        {
            BePort = bePort;
            IsUpdate = isUpdate;
            Password = password;
            Port = port;
            Resource = resource;
            SelfBuilt = selfBuilt;
            ServiceVip = serviceVip;
            UniqVpcId = uniqVpcId;
            UserName = userName;
        }
    }
}
