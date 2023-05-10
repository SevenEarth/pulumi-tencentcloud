// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr.Inputs
{

    public sealed class WebhookTriggerTriggerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// trigger rule.
        /// </summary>
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        /// <summary>
        /// trigger description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// enable trigger.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("eventTypes", required: true)]
        private InputList<string>? _eventTypes;

        /// <summary>
        /// trigger action.
        /// </summary>
        public InputList<string> EventTypes
        {
            get => _eventTypes ?? (_eventTypes = new InputList<string>());
            set => _eventTypes = value;
        }

        /// <summary>
        /// ID of the resource.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// trigger name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// the namespace Id to which the trigger belongs.
        /// </summary>
        [Input("namespaceId")]
        public Input<int>? NamespaceId { get; set; }

        [Input("targets", required: true)]
        private InputList<Inputs.WebhookTriggerTriggerTargetGetArgs>? _targets;

        /// <summary>
        /// trigger target.
        /// </summary>
        public InputList<Inputs.WebhookTriggerTriggerTargetGetArgs> Targets
        {
            get => _targets ?? (_targets = new InputList<Inputs.WebhookTriggerTriggerTargetGetArgs>());
            set => _targets = value;
        }

        public WebhookTriggerTriggerGetArgs()
        {
        }
    }
}
