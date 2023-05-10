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
    public sealed class LaunchTemplateVersionActionTimerExternalsStorageBlockAttr
    {
        /// <summary>
        /// Maximum capacity of local HDD storage.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// Minimum capacity of local HDD storage.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// Local HDD storage type. Value: LOCAL_PRO.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private LaunchTemplateVersionActionTimerExternalsStorageBlockAttr(
            int maxSize,

            int minSize,

            string type)
        {
            MaxSize = maxSize;
            MinSize = minSize;
            Type = type;
        }
    }
}
