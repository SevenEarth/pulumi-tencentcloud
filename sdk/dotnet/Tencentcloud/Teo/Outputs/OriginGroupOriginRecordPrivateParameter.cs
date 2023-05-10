// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class OriginGroupOriginRecordPrivateParameter
    {
        /// <summary>
        /// Parameter Name. Valid values:- AccessKeyId:Access Key ID.- SecretAccessKey:Secret Access Key.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Parameter value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private OriginGroupOriginRecordPrivateParameter(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}