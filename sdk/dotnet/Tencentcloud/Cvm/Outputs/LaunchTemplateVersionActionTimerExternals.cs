// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Outputs
{

    [OutputType]
    public sealed class LaunchTemplateVersionActionTimerExternals
    {
        /// <summary>
        /// Release address.
        /// </summary>
        public readonly bool? ReleaseAddress;
        /// <summary>
        /// Information on local HDD storage.
        /// </summary>
        public readonly Outputs.LaunchTemplateVersionActionTimerExternalsStorageBlockAttr? StorageBlockAttr;
        /// <summary>
        /// Not supported network.
        /// </summary>
        public readonly ImmutableArray<string> UnsupportNetworks;

        [OutputConstructor]
        private LaunchTemplateVersionActionTimerExternals(
            bool? releaseAddress,

            Outputs.LaunchTemplateVersionActionTimerExternalsStorageBlockAttr? storageBlockAttr,

            ImmutableArray<string> unsupportNetworks)
        {
            ReleaseAddress = releaseAddress;
            StorageBlockAttr = storageBlockAttr;
            UnsupportNetworks = unsupportNetworks;
        }
    }
}