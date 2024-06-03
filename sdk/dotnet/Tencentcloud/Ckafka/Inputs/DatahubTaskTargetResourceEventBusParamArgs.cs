// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTargetResourceEventBusParamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SCF function name.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// SCF namespace.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// SCF version and alias.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// instance id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Whether it is a self-built cluster.
        /// </summary>
        [Input("selfBuilt", required: true)]
        public Input<bool> SelfBuilt { get; set; } = null!;

        /// <summary>
        /// resource type. EB_COS/EB_ES/EB_CLS.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatahubTaskTargetResourceEventBusParamArgs()
        {
        }
        public static new DatahubTaskTargetResourceEventBusParamArgs Empty => new DatahubTaskTargetResourceEventBusParamArgs();
    }
}
