// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus.Inputs
{

    public sealed class JobConfigExpertModeConfigurationSlotSharingGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the SlotSharingGroupNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the SlotSharingGroupNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specification of the SlotSharingGroupNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("spec", required: true)]
        public Input<Inputs.JobConfigExpertModeConfigurationSlotSharingGroupSpecGetArgs> Spec { get; set; } = null!;

        public JobConfigExpertModeConfigurationSlotSharingGroupGetArgs()
        {
        }
        public static new JobConfigExpertModeConfigurationSlotSharingGroupGetArgs Empty => new JobConfigExpertModeConfigurationSlotSharingGroupGetArgs();
    }
}
