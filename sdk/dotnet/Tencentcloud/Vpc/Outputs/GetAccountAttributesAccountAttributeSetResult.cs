// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetAccountAttributesAccountAttributeSetResult
    {
        /// <summary>
        /// Attribute name.
        /// </summary>
        public readonly string AttributeName;
        /// <summary>
        /// Attribute values.
        /// </summary>
        public readonly ImmutableArray<string> AttributeValues;

        [OutputConstructor]
        private GetAccountAttributesAccountAttributeSetResult(
            string attributeName,

            ImmutableArray<string> attributeValues)
        {
            AttributeName = attributeName;
            AttributeValues = attributeValues;
        }
    }
}