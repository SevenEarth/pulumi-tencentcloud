// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ses.Inputs
{

    public sealed class TemplateTemplateContentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Html code after base64.
        /// </summary>
        [Input("html")]
        public Input<string>? Html { get; set; }

        /// <summary>
        /// Text content after base64.
        /// </summary>
        [Input("text")]
        public Input<string>? Text { get; set; }

        public TemplateTemplateContentGetArgs()
        {
        }
        public static new TemplateTemplateContentGetArgs Empty => new TemplateTemplateContentGetArgs();
    }
}
