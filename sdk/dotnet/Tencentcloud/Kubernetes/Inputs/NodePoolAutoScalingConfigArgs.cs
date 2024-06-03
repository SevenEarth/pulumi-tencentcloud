// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class NodePoolAutoScalingConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("backupInstanceTypes")]
        private InputList<string>? _backupInstanceTypes;

        /// <summary>
        /// Backup CVM instance types if specified instance type sold out or mismatch.
        /// </summary>
        public InputList<string> BackupInstanceTypes
        {
            get => _backupInstanceTypes ?? (_backupInstanceTypes = new InputList<string>());
            set => _backupInstanceTypes = value;
        }

        /// <summary>
        /// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
        /// </summary>
        [Input("bandwidthPackageId")]
        public Input<string>? BandwidthPackageId { get; set; }

        /// <summary>
        /// Name of cam role.
        /// </summary>
        [Input("camRoleName")]
        public Input<string>? CamRoleName { get; set; }

        [Input("dataDisks")]
        private InputList<Inputs.NodePoolAutoScalingConfigDataDiskArgs>? _dataDisks;

        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public InputList<Inputs.NodePoolAutoScalingConfigDataDiskArgs> DataDisks
        {
            get => _dataDisks ?? (_dataDisks = new InputList<Inputs.NodePoolAutoScalingConfigDataDiskArgs>());
            set => _dataDisks = value;
        }

        /// <summary>
        /// To specify whether to enable cloud monitor service. Default is TRUE.
        /// </summary>
        [Input("enhancedMonitorService")]
        public Input<bool>? EnhancedMonitorService { get; set; }

        /// <summary>
        /// To specify whether to enable cloud security service. Default is TRUE.
        /// </summary>
        [Input("enhancedSecurityService")]
        public Input<bool>? EnhancedSecurityService { get; set; }

        /// <summary>
        /// The hostname of the cloud server, dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows instances are not supported. Examples of other types (Linux, etc.): The character length is [2, 40], multiple periods are allowed, and there is a paragraph between the dots, and each paragraph is allowed to consist of letters (unlimited case), numbers and dashes (-). Pure numbers are not allowed. For usage, refer to `HostNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        [Input("hostName")]
        public Input<string>? HostName { get; set; }

        /// <summary>
        /// The style of the host name of the cloud server, the value range includes ORIGINAL and UNIQUE, and the default is ORIGINAL. For usage, refer to `HostNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        [Input("hostNameStyle")]
        public Input<string>? HostNameStyle { get; set; }

        /// <summary>
        /// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        /// </summary>
        [Input("instanceChargeTypePrepaidPeriod")]
        public Input<int>? InstanceChargeTypePrepaidPeriod { get; set; }

        /// <summary>
        /// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
        /// </summary>
        [Input("instanceChargeTypePrepaidRenewFlag")]
        public Input<string>? InstanceChargeTypePrepaidRenewFlag { get; set; }

        /// <summary>
        /// Instance name, no more than 60 characters. For usage, refer to `InstanceNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Specified types of CVM instance.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// Charge types for network traffic. Valid value: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// Max bandwidth of Internet access in Mbps. Default is `0`.
        /// </summary>
        [Input("internetMaxBandwidthOut")]
        public Input<int>? InternetMaxBandwidthOut { get; set; }

        [Input("keyIds")]
        private InputList<string>? _keyIds;

        /// <summary>
        /// ID list of keys.
        /// </summary>
        public InputList<string> KeyIds
        {
            get => _keyIds ?? (_keyIds = new InputList<string>());
            set => _keyIds = value;
        }

        [Input("orderlySecurityGroupIds")]
        private InputList<string>? _orderlySecurityGroupIds;

        /// <summary>
        /// Ordered security groups to which a CVM instance belongs.
        /// </summary>
        public InputList<string> OrderlySecurityGroupIds
        {
            get => _orderlySecurityGroupIds ?? (_orderlySecurityGroupIds = new InputList<string>());
            set => _orderlySecurityGroupIds = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Password to access.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specify whether to assign an Internet IP address.
        /// </summary>
        [Input("publicIpAssigned")]
        public Input<bool>? PublicIpAssigned { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The order of elements in this field cannot be guaranteed. Use `orderly_security_group_ids` instead. Security groups to which a CVM instance belongs.
        /// </summary>
        [Obsolete(@"The order of elements in this field cannot be guaranteed. Use `orderly_security_group_ids` instead.")]
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Type of spot instance, only support `one-time` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotInstanceType")]
        public Input<string>? SpotInstanceType { get; set; }

        /// <summary>
        /// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        [Input("spotMaxPrice")]
        public Input<string>? SpotMaxPrice { get; set; }

        /// <summary>
        /// Volume of system disk in GB. Default is `50`.
        /// </summary>
        [Input("systemDiskSize")]
        public Input<int>? SystemDiskSize { get; set; }

        /// <summary>
        /// Type of a CVM disk. Valid value: `LOCAL_BASIC`, `LOCAL_SSD`, `CLOUD_BASIC`, `CLOUD_PREMIUM`, `CLOUD_SSD`, `CLOUD_HSSD`, `CLOUD_TSSD` and `CLOUD_BSSD`. Default is `CLOUD_PREMIUM`.
        /// </summary>
        [Input("systemDiskType")]
        public Input<string>? SystemDiskType { get; set; }

        public NodePoolAutoScalingConfigArgs()
        {
        }
        public static new NodePoolAutoScalingConfigArgs Empty => new NodePoolAutoScalingConfigArgs();
    }
}
