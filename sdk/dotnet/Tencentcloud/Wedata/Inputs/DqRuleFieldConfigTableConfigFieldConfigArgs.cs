// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Wedata.Inputs
{

    public sealed class DqRuleFieldConfigTableConfigFieldConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Field typeNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("fieldDataType")]
        public Input<string>? FieldDataType { get; set; }

        /// <summary>
        /// Field keyNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("fieldKey")]
        public Input<string>? FieldKey { get; set; }

        /// <summary>
        /// Field valueNote: This field may return null, indicating that a valid value cannot be obtained.
        /// </summary>
        [Input("fieldValue")]
        public Input<string>? FieldValue { get; set; }

        public DqRuleFieldConfigTableConfigFieldConfigArgs()
        {
        }
        public static new DqRuleFieldConfigTableConfigFieldConfigArgs Empty => new DqRuleFieldConfigTableConfigFieldConfigArgs();
    }
}
