// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class ClusterExistInstanceArgs : global::Pulumi.ResourceArgs
    {
        [Input("desiredPodNumbers")]
        private InputList<int>? _desiredPodNumbers;

        /// <summary>
        /// Custom mode cluster, you can specify the number of pods for each node. corresponding to the existed_instances_para.instance_ids parameter.
        /// </summary>
        public InputList<int> DesiredPodNumbers
        {
            get => _desiredPodNumbers ?? (_desiredPodNumbers = new InputList<int>());
            set => _desiredPodNumbers = value;
        }

        /// <summary>
        /// Reinstallation parameters of an existing instance.
        /// </summary>
        [Input("instancesPara")]
        public Input<Inputs.ClusterExistInstanceInstancesParaArgs>? InstancesPara { get; set; }

        /// <summary>
        /// Role of existed node. value:MASTER_ETCD or WORKER.
        /// </summary>
        [Input("nodeRole")]
        public Input<string>? NodeRole { get; set; }

        public ClusterExistInstanceArgs()
        {
        }
        public static new ClusterExistInstanceArgs Empty => new ClusterExistInstanceArgs();
    }
}
