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

    public sealed class JobConfigExpertModeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Job graphNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("jobGraph")]
        public Input<Inputs.JobConfigExpertModeConfigurationJobGraphArgs>? JobGraph { get; set; }

        [Input("nodeConfigs")]
        private InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigArgs>? _nodeConfigs;

        /// <summary>
        /// Node configurationNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigArgs> NodeConfigs
        {
            get => _nodeConfigs ?? (_nodeConfigs = new InputList<Inputs.JobConfigExpertModeConfigurationNodeConfigArgs>());
            set => _nodeConfigs = value;
        }

        [Input("slotSharingGroups")]
        private InputList<Inputs.JobConfigExpertModeConfigurationSlotSharingGroupArgs>? _slotSharingGroups;

        /// <summary>
        /// Slot sharing groupsNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        public InputList<Inputs.JobConfigExpertModeConfigurationSlotSharingGroupArgs> SlotSharingGroups
        {
            get => _slotSharingGroups ?? (_slotSharingGroups = new InputList<Inputs.JobConfigExpertModeConfigurationSlotSharingGroupArgs>());
            set => _slotSharingGroups = value;
        }

        public JobConfigExpertModeConfigurationArgs()
        {
        }
        public static new JobConfigExpertModeConfigurationArgs Empty => new JobConfigExpertModeConfigurationArgs();
    }
}
