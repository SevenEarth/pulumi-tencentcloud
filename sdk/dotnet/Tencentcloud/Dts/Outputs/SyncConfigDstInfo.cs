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
    public sealed class SyncConfigDstInfo
    {
        /// <summary>
        /// The account to which the instance belongs. This field is required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Account;
        /// <summary>
        /// The account to which the resource belongs is empty or self (represents resources within this account), other (represents cross-account resources). Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? AccountMode;
        /// <summary>
        /// The role during cross-account synchronization, only [a-zA-Z0-9-_]+ is allowed, if it is a cross-account instance, this field is required. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? AccountRole;
        /// <summary>
        /// Cloud networking ID, which is required for the cloud networking access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? CcnId;
        /// <summary>
        /// CVM instance short ID, which is the same as the instance ID displayed on the cloud server console page. If it is a self-built instance of CVM, this field needs to be passed. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? CvmInstanceId;
        /// <summary>
        /// Database kernel type, used to distinguish different kernels in tdsql: percona, mariadb, mysql. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? DbKernel;
        /// <summary>
        /// Database name, when the database is cdwpg, it needs to be provided. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? DbName;
        /// <summary>
        /// Whether to use encrypted transmission, UnEncrypted means not to use encrypted transmission, Encrypted means to use encrypted transmission, the default is UnEncrypted. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? EncryptConn;
        /// <summary>
        /// Database version, valid only when the instance is an RDS instance, ignored by other instances, the format is: 5.6 or 5.7, the default is 5.6. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? EngineVersion;
        /// <summary>
        /// Database instance id. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? InstanceId;
        /// <summary>
        /// The IP address of the instance, which is required when the access type is non-cdb. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// Password, required for instances that require username and password authentication for access. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Instance port, this item is required when the access type is non-cdb. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The english name of region. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// The node type of tdsql mysql version, the enumeration value is proxy, set. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Role;
        /// <summary>
        /// External role id. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? RoleExternalId;
        /// <summary>
        /// The subnet ID under the private network, this item is required for the private network, leased line, and VPN access methods. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Cloud vendor type, when the instance is an RDS instance, fill in aliyun, in other cases fill in others, the default is others. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? Supplier;
        /// <summary>
        /// Temporary key Id, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? TmpSecretId;
        /// <summary>
        /// Temporary key Key, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? TmpSecretKey;
        /// <summary>
        /// Temporary Token, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? TmpToken;
        /// <summary>
        /// Leased line gateway ID, which is required for the leased line access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? UniqDcgId;
        /// <summary>
        /// VPN gateway ID, which is required for the VPN access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? UniqVpnGwId;
        /// <summary>
        /// Username, required for instances that require username and password authentication for access. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? User;
        /// <summary>
        /// Private network ID, which is required for access methods of private network, leased line, and VPN. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private SyncConfigDstInfo(
            string? account,

            string? accountMode,

            string? accountRole,

            string? ccnId,

            string? cvmInstanceId,

            string? dbKernel,

            string? dbName,

            string? encryptConn,

            string? engineVersion,

            string? instanceId,

            string? ip,

            string? password,

            int? port,

            string? region,

            string? role,

            string? roleExternalId,

            string? subnetId,

            string? supplier,

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
            CcnId = ccnId;
            CvmInstanceId = cvmInstanceId;
            DbKernel = dbKernel;
            DbName = dbName;
            EncryptConn = encryptConn;
            EngineVersion = engineVersion;
            InstanceId = instanceId;
            Ip = ip;
            Password = password;
            Port = port;
            Region = region;
            Role = role;
            RoleExternalId = roleExternalId;
            SubnetId = subnetId;
            Supplier = supplier;
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