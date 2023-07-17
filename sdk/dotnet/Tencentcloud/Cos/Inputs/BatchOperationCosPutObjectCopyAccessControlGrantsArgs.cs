// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BatchOperationCosPutObjectCopyAccessControlGrantsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// User name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// User ID (UIN) in qcs format. For example: qcs::cam::uin/100000000001:uin/100000000001.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        /// <summary>
        /// Specify a permission to be granted. Enumerated value: READ,WRITE,FULL_CONTROL.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// Specifies the type of Identifier. Currently, only user ID is supported. Enumerated value: ID.
        /// </summary>
        [Input("typeIdentifier", required: true)]
        public Input<string> TypeIdentifier { get; set; } = null!;

        public BatchOperationCosPutObjectCopyAccessControlGrantsArgs()
        {
        }
    }
}
