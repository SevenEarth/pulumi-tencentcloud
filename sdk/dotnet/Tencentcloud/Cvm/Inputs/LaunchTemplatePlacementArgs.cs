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

    public sealed class LaunchTemplatePlacementArgs : Pulumi.ResourceArgs
    {
        [Input("hostIds")]
        private InputList<string>? _hostIds;

        /// <summary>
        /// The CDH ID list of the instance(input).
        /// </summary>
        public InputList<string> HostIds
        {
            get => _hostIds ?? (_hostIds = new InputList<string>());
            set => _hostIds = value;
        }

        [Input("hostIps")]
        private InputList<string>? _hostIps;

        /// <summary>
        /// Specify the host machine ip.
        /// </summary>
        public InputList<string> HostIps
        {
            get => _hostIps ?? (_hostIps = new InputList<string>());
            set => _hostIps = value;
        }

        /// <summary>
        /// The project ID of the instance.
        /// </summary>
        [Input("projectId")]
        public Input<int>? ProjectId { get; set; }

        /// <summary>
        /// The available zone ID of the instance.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public LaunchTemplatePlacementArgs()
        {
        }
    }
}
