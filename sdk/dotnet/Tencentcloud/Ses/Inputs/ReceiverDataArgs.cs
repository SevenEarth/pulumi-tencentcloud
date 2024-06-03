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

    public sealed class ReceiverDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Recipient email addresses.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Variable parameters in the template, please use json.dump to format the JSON object as a string type. The object is a set of key-value pairs, where each key represents a variable in the template, and the variables in the template are represented by {{key}}, and the corresponding values will be replaced with {{value}} when sent.Note: Parameter values cannot be complex data such as HTML. The total length of TemplateData (the entire JSON structure) should be less than 800 bytes.
        /// </summary>
        [Input("templateData")]
        public Input<string>? TemplateData { get; set; }

        public ReceiverDataArgs()
        {
        }
        public static new ReceiverDataArgs Empty => new ReceiverDataArgs();
    }
}
