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

    public sealed class AiRecognitionTemplateAsrWordsConfigureArgs : Pulumi.ResourceArgs
    {
        [Input("labelSets")]
        private InputList<string>? _labelSets;

        /// <summary>
        /// Keyword filter label, specify the label of the keyword to be returned. If not filled or empty, all results will be returned.The maximum number of tags is 10, and the length of each tag is up to 16 characters.
        /// </summary>
        public InputList<string> LabelSets
        {
            get => _labelSets ?? (_labelSets = new InputList<string>());
            set => _labelSets = value;
        }

        /// <summary>
        /// Asr word recognition task switch, optional value:ON/OFF.
        /// </summary>
        [Input("switch", required: true)]
        public Input<string> Switch { get; set; } = null!;

        public AiRecognitionTemplateAsrWordsConfigureArgs()
        {
        }
    }
}
