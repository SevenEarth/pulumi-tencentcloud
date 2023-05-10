// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class MigrateJobSrcInfoInfo
    {
        /// <summary>
        /// Account.
        /// </summary>
        public readonly string? Account;
        /// <summary>
        /// AccountMode.
        /// </summary>
        public readonly string? AccountMode;
        /// <summary>
        /// AccountRole.
        /// </summary>
        public readonly string? AccountRole;
        /// <summary>
        /// CcnGwId.
        /// </summary>
        public readonly string? CcnGwId;
        /// <summary>
        /// CvmInstanceId.
        /// </summary>
        public readonly string? CvmInstanceId;
        /// <summary>
        /// DbKernel.
        /// </summary>
        public readonly string? DbKernel;
        /// <summary>
        /// EngineVersion.
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// Host.
        /// </summary>
        public readonly string? Host;
        /// <summary>
        /// InstanceId.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// Password.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Port.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Role.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// SubnetId.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// TmpSecretId.
        /// </summary>
        public readonly string? TmpSecretId;
        /// <summary>
        /// TmpSecretKey.
        /// </summary>
        public readonly string? TmpSecretKey;
        /// <summary>
        /// TmpToken.
        /// </summary>
        public readonly string? TmpToken;
        /// <summary>
        /// UniqDcgId.
        /// </summary>
        public readonly string? UniqDcgId;
        /// <summary>
        /// UniqVpnGwId.
        /// </summary>
        public readonly string? UniqVpnGwId;
        /// <summary>
        /// User.
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// VpcId.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private MigrateJobSrcInfoInfo(
            string? account,

            string? accountMode,

            string? accountRole,

            string? ccnGwId,

            string? cvmInstanceId,

            string? dbKernel,

            string? engineVersion,

            string? host,

            string? instanceId,

            string? password,

            int? port,

            string? role,

            string? subnetId,

            string? tmpSecretId,

            string? tmpSecretKey,

            string? tmpToken,

            string? uniqDcgId,

            string? uniqVpnGwId,

            string? user,

            string? vpcId)
        {
            Account = account;
            AccountMode = accountMode;
            AccountRole = accountRole;
            CcnGwId = ccnGwId;
            CvmInstanceId = cvmInstanceId;
            DbKernel = dbKernel;
            EngineVersion = engineVersion;
            Host = host;
            InstanceId = instanceId;
            Password = password;
            Port = port;
            Role = role;
            SubnetId = subnetId;
            TmpSecretId = tmpSecretId;
            TmpSecretKey = tmpSecretKey;
            TmpToken = tmpToken;
            UniqDcgId = uniqDcgId;
            UniqVpnGwId = uniqVpnGwId;
            User = user;
            VpcId = vpcId;
        }
    }
}
