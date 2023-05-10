// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Outputs
{

    [OutputType]
    public sealed class NodePoolAutoScalingConfig
    {
        /// <summary>
        /// Backup CVM instance types if specified instance type sold out or mismatch.
        /// </summary>
        public readonly ImmutableArray<string> BackupInstanceTypes;
        /// <summary>
        /// bandwidth package id. if user is standard user, then the bandwidth_package_id is needed, or default has bandwidth_package_id.
        /// </summary>
        public readonly string? BandwidthPackageId;
        /// <summary>
        /// Name of cam role.
        /// </summary>
        public readonly string? CamRoleName;
        /// <summary>
        /// Configurations of data disk.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodePoolAutoScalingConfigDataDisk> DataDisks;
        /// <summary>
        /// To specify whether to enable cloud monitor service. Default is TRUE.
        /// </summary>
        public readonly bool? EnhancedMonitorService;
        /// <summary>
        /// To specify whether to enable cloud security service. Default is TRUE.
        /// </summary>
        public readonly bool? EnhancedSecurityService;
        /// <summary>
        /// The hostname of the cloud server, dot (.) and dash (-) cannot be used as the first and last characters of HostName and cannot be used consecutively. Windows instances are not supported. Examples of other types (Linux, etc.): The character length is [2, 40], multiple periods are allowed, and there is a paragraph between the dots, and each paragraph is allowed to consist of letters (unlimited case), numbers and dashes (-). Pure numbers are not allowed. For usage, refer to `HostNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        public readonly string? HostName;
        /// <summary>
        /// The style of the host name of the cloud server, the value range includes ORIGINAL and UNIQUE, and the default is ORIGINAL. For usage, refer to `HostNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        public readonly string? HostNameStyle;
        /// <summary>
        /// Charge type of instance. Valid values are `PREPAID`, `POSTPAID_BY_HOUR`, `SPOTPAID`. The default is `POSTPAID_BY_HOUR`. NOTE: `SPOTPAID` instance must set `spot_instance_type` and `spot_max_price` at the same time.
        /// </summary>
        public readonly string? InstanceChargeType;
        /// <summary>
        /// The tenancy (in month) of the prepaid instance, NOTE: it only works when instance_charge_type is set to `PREPAID`. Valid values are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`, `11`, `12`, `24`, `36`.
        /// </summary>
        public readonly int? InstanceChargeTypePrepaidPeriod;
        /// <summary>
        /// Auto renewal flag. Valid values: `NOTIFY_AND_AUTO_RENEW`: notify upon expiration and renew automatically, `NOTIFY_AND_MANUAL_RENEW`: notify upon expiration but do not renew automatically, `DISABLE_NOTIFY_AND_MANUAL_RENEW`: neither notify upon expiration nor renew automatically. Default value: `NOTIFY_AND_MANUAL_RENEW`. If this parameter is specified as `NOTIFY_AND_AUTO_RENEW`, the instance will be automatically renewed on a monthly basis if the account balance is sufficient. NOTE: it only works when instance_charge_type is set to `PREPAID`.
        /// </summary>
        public readonly string? InstanceChargeTypePrepaidRenewFlag;
        /// <summary>
        /// Instance name, no more than 60 characters. For usage, refer to `InstanceNameSettings` in https://www.tencentcloud.com/document/product/377/31001.
        /// </summary>
        public readonly string? InstanceName;
        /// <summary>
        /// Specified types of CVM instance.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Charge types for network traffic. Valid value: `BANDWIDTH_PREPAID`, `TRAFFIC_POSTPAID_BY_HOUR`, `TRAFFIC_POSTPAID_BY_HOUR` and `BANDWIDTH_PACKAGE`.
        /// </summary>
        public readonly string? InternetChargeType;
        /// <summary>
        /// Max bandwidth of Internet access in Mbps. Default is `0`.
        /// </summary>
        public readonly int? InternetMaxBandwidthOut;
        /// <summary>
        /// ID list of keys.
        /// </summary>
        public readonly ImmutableArray<string> KeyIds;
        /// <summary>
        /// Password to access.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Specify whether to assign an Internet IP address.
        /// </summary>
        public readonly bool? PublicIpAssigned;
        /// <summary>
        /// Security groups to which a CVM instance belongs.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// Type of spot instance, only support `one-time` now. Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        public readonly string? SpotInstanceType;
        /// <summary>
        /// Max price of a spot instance, is the format of decimal string, for example "0.50". Note: it only works when instance_charge_type is set to `SPOTPAID`.
        /// </summary>
        public readonly string? SpotMaxPrice;
        /// <summary>
        /// Volume of system disk in GB. Default is `50`.
        /// </summary>
        public readonly int? SystemDiskSize;
        /// <summary>
        /// Type of a CVM disk. Valid value: `CLOUD_PREMIUM` and `CLOUD_SSD`. Default is `CLOUD_PREMIUM`.
        /// </summary>
        public readonly string? SystemDiskType;

        [OutputConstructor]
        private NodePoolAutoScalingConfig(
            ImmutableArray<string> backupInstanceTypes,

            string? bandwidthPackageId,

            string? camRoleName,

            ImmutableArray<Outputs.NodePoolAutoScalingConfigDataDisk> dataDisks,

            bool? enhancedMonitorService,

            bool? enhancedSecurityService,

            string? hostName,

            string? hostNameStyle,

            string? instanceChargeType,

            int? instanceChargeTypePrepaidPeriod,

            string? instanceChargeTypePrepaidRenewFlag,

            string? instanceName,

            string instanceType,

            string? internetChargeType,

            int? internetMaxBandwidthOut,

            ImmutableArray<string> keyIds,

            string? password,

            bool? publicIpAssigned,

            ImmutableArray<string> securityGroupIds,

            string? spotInstanceType,

            string? spotMaxPrice,

            int? systemDiskSize,

            string? systemDiskType)
        {
            BackupInstanceTypes = backupInstanceTypes;
            BandwidthPackageId = bandwidthPackageId;
            CamRoleName = camRoleName;
            DataDisks = dataDisks;
            EnhancedMonitorService = enhancedMonitorService;
            EnhancedSecurityService = enhancedSecurityService;
            HostName = hostName;
            HostNameStyle = hostNameStyle;
            InstanceChargeType = instanceChargeType;
            InstanceChargeTypePrepaidPeriod = instanceChargeTypePrepaidPeriod;
            InstanceChargeTypePrepaidRenewFlag = instanceChargeTypePrepaidRenewFlag;
            InstanceName = instanceName;
            InstanceType = instanceType;
            InternetChargeType = internetChargeType;
            InternetMaxBandwidthOut = internetMaxBandwidthOut;
            KeyIds = keyIds;
            Password = password;
            PublicIpAssigned = publicIpAssigned;
            SecurityGroupIds = securityGroupIds;
            SpotInstanceType = spotInstanceType;
            SpotMaxPrice = spotMaxPrice;
            SystemDiskSize = systemDiskSize;
            SystemDiskType = systemDiskType;
        }
    }
}
