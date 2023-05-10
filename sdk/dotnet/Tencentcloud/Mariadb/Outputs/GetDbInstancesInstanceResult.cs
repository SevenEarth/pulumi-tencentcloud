// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mariadb.Outputs
{

    [OutputType]
    public sealed class GetDbInstancesInstanceResult
    {
        /// <summary>
        /// db version id.
        /// </summary>
        public readonly string DbVersionId;
        /// <summary>
        /// instance id.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// instance name.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// meory of instance.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// project id.
        /// </summary>
        public readonly int ProjectId;
        /// <summary>
        /// region.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// resource tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDbInstancesInstanceResourceTagResult> ResourceTags;
        /// <summary>
        /// storage of instance.
        /// </summary>
        public readonly int Storage;
        /// <summary>
        /// subnet id.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// vpc id.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// available zone.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetDbInstancesInstanceResult(
            string dbVersionId,

            string instanceId,

            string instanceName,

            int memory,

            int projectId,

            string region,

            ImmutableArray<Outputs.GetDbInstancesInstanceResourceTagResult> resourceTags,

            int storage,

            string subnetId,

            string vpcId,

            string zone)
        {
            DbVersionId = dbVersionId;
            InstanceId = instanceId;
            InstanceName = instanceName;
            Memory = memory;
            ProjectId = projectId;
            Region = region;
            ResourceTags = resourceTags;
            Storage = storage;
            SubnetId = subnetId;
            VpcId = vpcId;
            Zone = zone;
        }
    }
}
