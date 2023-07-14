// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Inputs
{

    public sealed class SyncConfigSrcInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account to which the instance belongs. This field is required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("account")]
        public Input<string>? Account { get; set; }

        /// <summary>
        /// The account to which the resource belongs is empty or self (represents resources within this account), other (represents cross-account resources). Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("accountMode")]
        public Input<string>? AccountMode { get; set; }

        /// <summary>
        /// The role during cross-account synchronization, only [a-zA-Z0-9-_]+ is allowed, if it is a cross-account instance, this field is required. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("accountRole")]
        public Input<string>? AccountRole { get; set; }

        /// <summary>
        /// Cloud networking ID, which is required for the cloud networking access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("ccnId")]
        public Input<string>? CcnId { get; set; }

        /// <summary>
        /// CVM instance short ID, which is the same as the instance ID displayed on the cloud server console page. If it is a self-built instance of CVM, this field needs to be passed. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("cvmInstanceId")]
        public Input<string>? CvmInstanceId { get; set; }

        /// <summary>
        /// Database kernel type, used to distinguish different kernels in tdsql: percona, mariadb, mysql. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("dbKernel")]
        public Input<string>? DbKernel { get; set; }

        /// <summary>
        /// Database name, when the database is cdwpg, it needs to be provided. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// Whether to use encrypted transmission, UnEncrypted means not to use encrypted transmission, Encrypted means to use encrypted transmission, the default is UnEncrypted. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("encryptConn")]
        public Input<string>? EncryptConn { get; set; }

        /// <summary>
        /// Database version, valid only when the instance is an RDS instance, ignored by other instances, the format is: 5.6 or 5.7, the default is 5.6. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Database instance id. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The IP address of the instance, which is required when the access type is non-cdb. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Password, required for instances that require username and password authentication for access. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Instance port, this item is required when the access type is non-cdb. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The english name of region. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The node type of tdsql mysql version, the enumeration value is proxy, set. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// External role id. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("roleExternalId")]
        public Input<string>? RoleExternalId { get; set; }

        /// <summary>
        /// The subnet ID under the private network, this item is required for the private network, leased line, and VPN access methods. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// Cloud vendor type, when the instance is an RDS instance, fill in aliyun, in other cases fill in others, the default is others. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("supplier")]
        public Input<string>? Supplier { get; set; }

        /// <summary>
        /// Temporary key Id, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("tmpSecretId")]
        public Input<string>? TmpSecretId { get; set; }

        /// <summary>
        /// Temporary key Key, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("tmpSecretKey")]
        public Input<string>? TmpSecretKey { get; set; }

        /// <summary>
        /// Temporary Token, required if it is a cross-account instance. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("tmpToken")]
        public Input<string>? TmpToken { get; set; }

        /// <summary>
        /// Leased line gateway ID, which is required for the leased line access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("uniqDcgId")]
        public Input<string>? UniqDcgId { get; set; }

        /// <summary>
        /// VPN gateway ID, which is required for the VPN access type. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("uniqVpnGwId")]
        public Input<string>? UniqVpnGwId { get; set; }

        /// <summary>
        /// Username, required for instances that require username and password authentication for access. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        /// <summary>
        /// Private network ID, which is required for access methods of private network, leased line, and VPN. Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public SyncConfigSrcInfoArgs()
        {
        }
    }
}
