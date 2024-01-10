// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class FlowInputGroupRtspPullSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("sourceAddresses", required: true)]
        private InputList<Inputs.FlowInputGroupRtspPullSettingsSourceAddressArgs>? _sourceAddresses;

        /// <summary>
        /// The source site address of the RTSP source site, there can only be one.
        /// </summary>
        public InputList<Inputs.FlowInputGroupRtspPullSettingsSourceAddressArgs> SourceAddresses
        {
            get => _sourceAddresses ?? (_sourceAddresses = new InputList<Inputs.FlowInputGroupRtspPullSettingsSourceAddressArgs>());
            set => _sourceAddresses = value;
        }

        public FlowInputGroupRtspPullSettingsArgs()
        {
        }
    }
}
