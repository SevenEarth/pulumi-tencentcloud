// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class ClusterOperationInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("addInstances")]
        private InputList<Inputs.ClusterOperationInfoAddInstanceArgs>? _addInstances;

        /// <summary>
        /// Add the control information of the instance button.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoAddInstanceArgs> AddInstances
        {
            get => _addInstances ?? (_addInstances = new InputList<Inputs.ClusterOperationInfoAddInstanceArgs>());
            set => _addInstances = value;
        }

        [Input("destroys")]
        private InputList<Inputs.ClusterOperationInfoDestroyArgs>? _destroys;

        /// <summary>
        /// Destroy the control information of the machine.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoDestroyArgs> Destroys
        {
            get => _destroys ?? (_destroys = new InputList<Inputs.ClusterOperationInfoDestroyArgs>());
            set => _destroys = value;
        }

        [Input("inits")]
        private InputList<Inputs.ClusterOperationInfoInitArgs>? _inits;

        /// <summary>
        /// Initialize the control information of the button.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoInitArgs> Inits
        {
            get => _inits ?? (_inits = new InputList<Inputs.ClusterOperationInfoInitArgs>());
            set => _inits = value;
        }

        public ClusterOperationInfoArgs()
        {
        }
        public static new ClusterOperationInfoArgs Empty => new ClusterOperationInfoArgs();
    }
}
