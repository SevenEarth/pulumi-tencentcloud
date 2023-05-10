// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm.Inputs
{

    public sealed class LaunchTemplateVersionActionTimerExternalsStorageBlockAttrArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum capacity of local HDD storage.
        /// </summary>
        [Input("maxSize", required: true)]
        public Input<int> MaxSize { get; set; } = null!;

        /// <summary>
        /// Minimum capacity of local HDD storage.
        /// </summary>
        [Input("minSize", required: true)]
        public Input<int> MinSize { get; set; } = null!;

        /// <summary>
        /// Local HDD storage type. Value: LOCAL_PRO.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public LaunchTemplateVersionActionTimerExternalsStorageBlockAttrArgs()
        {
        }
    }
}
