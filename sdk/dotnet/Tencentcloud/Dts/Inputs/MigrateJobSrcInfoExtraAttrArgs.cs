// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Inputs
{

    public sealed class MigrateJobSrcInfoExtraAttrArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public MigrateJobSrcInfoExtraAttrArgs()
        {
        }
    }
}
