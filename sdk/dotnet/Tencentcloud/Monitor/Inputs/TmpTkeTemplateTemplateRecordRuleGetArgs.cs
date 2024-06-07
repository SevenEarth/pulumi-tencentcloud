// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class TmpTkeTemplateTemplateRecordRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Config.
        /// </summary>
        [Input("config", required: true)]
        public Input<string> Config { get; set; } = null!;

        /// <summary>
        /// Name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Used for the argument, if the configuration comes to the template, the template id.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public TmpTkeTemplateTemplateRecordRuleGetArgs()
        {
        }
        public static new TmpTkeTemplateTemplateRecordRuleGetArgs Empty => new TmpTkeTemplateTemplateRecordRuleGetArgs();
    }
}
