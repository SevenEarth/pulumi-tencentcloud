// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Outputs
{

    [OutputType]
    public sealed class InstancesAttachmentInstanceAdvancedSettings
    {
        /// <summary>
        /// Dockerd --graph specifies the value, default is /var/lib/docker Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string DockerGraphPath;
        /// <summary>
        /// Data disk mount point, data disks are not mounted by default. Data disks with formatted ext3, ext4, xfs file systems will be mounted directly, other file systems or unformatted data disks will be automatically formatted as ext4 and mounted. Please back up your data! This setting does not take effect for cloud servers with no data disks or multiple data disks. Note: This field may return null, indicating that no valid values can be obtained.
        /// </summary>
        public readonly string MountTarget;

        [OutputConstructor]
        private InstancesAttachmentInstanceAdvancedSettings(
            string dockerGraphPath,

            string mountTarget)
        {
            DockerGraphPath = dockerGraphPath;
            MountTarget = mountTarget;
        }
    }
}
