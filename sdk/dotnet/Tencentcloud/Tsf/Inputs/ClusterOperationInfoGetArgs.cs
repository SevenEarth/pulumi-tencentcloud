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

    public sealed class ClusterOperationInfoGetArgs : Pulumi.ResourceArgs
    {
        [Input("addInstances")]
        private InputList<Inputs.ClusterOperationInfoAddInstanceGetArgs>? _addInstances;

        /// <summary>
        /// Add the control information of the instance button.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoAddInstanceGetArgs> AddInstances
        {
            get => _addInstances ?? (_addInstances = new InputList<Inputs.ClusterOperationInfoAddInstanceGetArgs>());
            set => _addInstances = value;
        }

        [Input("destroys")]
        private InputList<Inputs.ClusterOperationInfoDestroyGetArgs>? _destroys;

        /// <summary>
        /// Destroy the control information of the machine.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoDestroyGetArgs> Destroys
        {
            get => _destroys ?? (_destroys = new InputList<Inputs.ClusterOperationInfoDestroyGetArgs>());
            set => _destroys = value;
        }

        [Input("inits")]
        private InputList<Inputs.ClusterOperationInfoInitGetArgs>? _inits;

        /// <summary>
        /// Initialize the control information of the button.
        /// </summary>
        public InputList<Inputs.ClusterOperationInfoInitGetArgs> Inits
        {
            get => _inits ?? (_inits = new InputList<Inputs.ClusterOperationInfoInitGetArgs>());
            set => _inits = value;
        }

        public ClusterOperationInfoGetArgs()
        {
        }
    }
}
