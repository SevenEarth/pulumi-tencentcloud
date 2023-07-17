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
    public sealed class GetDatahubTaskTaskListTargetResourceDtsParamResult
    {
        /// <summary>
        /// Dts consumer group Id.
        /// </summary>
        public readonly string GroupId;
        /// <summary>
        /// Dts consumer group passwd.
        /// </summary>
        public readonly string GroupPassword;
        /// <summary>
        /// Dts account.
        /// </summary>
        public readonly string GroupUser;
        /// <summary>
        /// Mongo DB connection ip.
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// MongoDB connection port.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Resource.
        /// </summary>
        public readonly string Resource;
        /// <summary>
        /// Topic name, use `,` when more than 1 topic.
        /// </summary>
        public readonly string Topic;
        /// <summary>
        /// False to synchronize the original data, true to synchronize the parsed json format data, the default is true.
        /// </summary>
        public readonly bool TranSql;

        [OutputConstructor]
        private GetDatahubTaskTaskListTargetResourceDtsParamResult(
            string groupId,

            string groupPassword,

            string groupUser,

            string ip,

            int port,

            string resource,

            string topic,

            bool tranSql)
        {
            GroupId = groupId;
            GroupPassword = groupPassword;
            GroupUser = groupUser;
            Ip = ip;
            Port = port;
            Resource = resource;
            Topic = topic;
            TranSql = tranSql;
        }
    }
}
