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
    public sealed class LaunchTemplateActionTimerExternals
    {
        /// <summary>
        /// Release address.
        /// </summary>
        public readonly bool? ReleaseAddress;
        /// <summary>
        /// HDD local storage attributes.
        /// </summary>
        public readonly Outputs.LaunchTemplateActionTimerExternalsStorageBlockAttr? StorageBlockAttr;
        /// <summary>
        /// Unsupported network type.
        /// </summary>
        public readonly ImmutableArray<string> UnsupportNetworks;

        [OutputConstructor]
        private LaunchTemplateActionTimerExternals(
            bool? releaseAddress,

            Outputs.LaunchTemplateActionTimerExternalsStorageBlockAttr? storageBlockAttr,

            ImmutableArray<string> unsupportNetworks)
        {
            ReleaseAddress = releaseAddress;
            StorageBlockAttr = storageBlockAttr;
            UnsupportNetworks = unsupportNetworks;
        }
    }
}
