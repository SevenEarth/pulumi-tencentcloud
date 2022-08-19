// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Eks.Inputs
{

    public sealed class ContainerInstanceInitContainerVolumeMountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Volume mount propagation.
        /// </summary>
        [Input("mountPropagation")]
        public Input<string>? MountPropagation { get; set; }

        /// <summary>
        /// Volume name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Volume mount path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Whether the volume is read-only.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// Volume mount sub-path.
        /// </summary>
        [Input("subPath")]
        public Input<string>? SubPath { get; set; }

        /// <summary>
        /// Volume mount sub-path expression.
        /// </summary>
        [Input("subPathExpr")]
        public Input<string>? SubPathExpr { get; set; }

        public ContainerInstanceInitContainerVolumeMountGetArgs()
        {
        }
    }
}
