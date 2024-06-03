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

    public sealed class ClusterClusterExtraArgsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("kubeApiservers")]
        private InputList<string>? _kubeApiservers;

        /// <summary>
        /// The customized parameters for kube-apiserver.
        /// </summary>
        public InputList<string> KubeApiservers
        {
            get => _kubeApiservers ?? (_kubeApiservers = new InputList<string>());
            set => _kubeApiservers = value;
        }

        [Input("kubeControllerManagers")]
        private InputList<string>? _kubeControllerManagers;

        /// <summary>
        /// The customized parameters for kube-controller-manager.
        /// </summary>
        public InputList<string> KubeControllerManagers
        {
            get => _kubeControllerManagers ?? (_kubeControllerManagers = new InputList<string>());
            set => _kubeControllerManagers = value;
        }

        [Input("kubeSchedulers")]
        private InputList<string>? _kubeSchedulers;

        /// <summary>
        /// The customized parameters for kube-scheduler.
        /// </summary>
        public InputList<string> KubeSchedulers
        {
            get => _kubeSchedulers ?? (_kubeSchedulers = new InputList<string>());
            set => _kubeSchedulers = value;
        }

        public ClusterClusterExtraArgsGetArgs()
        {
        }
        public static new ClusterClusterExtraArgsGetArgs Empty => new ClusterClusterExtraArgsGetArgs();
    }
}
