// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu.Inputs
{

    public sealed class CcPolicyV2CcPrecisionPolicyPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration item types, currently only support value.
        /// </summary>
        [Input("fieldName", required: true)]
        public Input<string> FieldName { get; set; } = null!;

        /// <summary>
        /// Configuration fields with the desirable values cgi, ua, cookie, referer, accept, srcip.
        /// </summary>
        [Input("fieldType", required: true)]
        public Input<string> FieldType { get; set; } = null!;

        /// <summary>
        /// Configure the value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Configure the item-value comparison mode, which can be taken as the value of evaluate, not_equal, include.
        /// </summary>
        [Input("valueOperator", required: true)]
        public Input<string> ValueOperator { get; set; } = null!;

        public CcPolicyV2CcPrecisionPolicyPolicyArgs()
        {
        }
        public static new CcPolicyV2CcPrecisionPolicyPolicyArgs Empty => new CcPolicyV2CcPrecisionPolicyPolicyArgs();
    }
}
