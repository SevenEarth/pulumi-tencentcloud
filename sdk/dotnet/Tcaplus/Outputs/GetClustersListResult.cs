// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Tcaplus.Outputs
{

    [OutputType]
    public sealed class GetClustersListResult
    {
        /// <summary>
        /// Access id of the TcaplusDB cluster.For TcaplusDB SDK connect.
        /// </summary>
        public readonly string ApiAccessId;
        /// <summary>
        /// Access ip of the TcaplusDB cluster.For TcaplusDB SDK connect.
        /// </summary>
        public readonly string ApiAccessIp;
        /// <summary>
        /// Access port of the TcaplusDB cluster.For TcaplusDB SDK connect.
        /// </summary>
        public readonly int ApiAccessPort;
        /// <summary>
        /// ID of the TcaplusDB cluster to be query.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// Name of the TcaplusDB cluster to be query.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// Create time of the TcaplusDB cluster.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// IDL type of the TcaplusDB cluster.
        /// </summary>
        public readonly string IdlType;
        /// <summary>
        /// Network type of the TcaplusDB cluster.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// Expiration time of the old password. If `password_status` is `unmodifiable`, it means the old password has not yet expired.
        /// </summary>
        public readonly string OldPasswordExpireTime;
        /// <summary>
        /// Access password of the TcaplusDB cluster.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Password status of the TcaplusDB cluster. Valid values: `unmodifiable`, `modifiable`. `unmodifiable` means the password can not be changed in this moment; `modifiable` means the password can be changed in this moment.
        /// </summary>
        public readonly string PasswordStatus;
        /// <summary>
        /// Subnet ID of the TcaplusDB cluster.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// VPC ID of the TcaplusDB cluster.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetClustersListResult(
            string apiAccessId,

            string apiAccessIp,

            int apiAccessPort,

            string clusterId,

            string clusterName,

            string createTime,

            string idlType,

            string networkType,

            string oldPasswordExpireTime,

            string password,

            string passwordStatus,

            string subnetId,

            string vpcId)
        {
            ApiAccessId = apiAccessId;
            ApiAccessIp = apiAccessIp;
            ApiAccessPort = apiAccessPort;
            ClusterId = clusterId;
            ClusterName = clusterName;
            CreateTime = createTime;
            IdlType = idlType;
            NetworkType = networkType;
            OldPasswordExpireTime = oldPasswordExpireTime;
            Password = password;
            PasswordStatus = passwordStatus;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
