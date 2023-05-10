// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes
{
    /// <summary>
    /// Provide a resource to create a KubernetesClusterEndpoint. This resource allows you to create an empty cluster first without any workers. Only all attached node depends create complete, cluster endpoint will finally be enabled.
    /// 
    /// &gt; **NOTE:** Recommend using `depends_on` to make sure endpoint create after node pools or workers does.
    /// 
    /// ## Import
    /// 
    /// KubernetesClusterEndpoint instance can be imported by passing cluster id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint test cluster-id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint")]
    public partial class ClusterEndpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The certificate used for access.
        /// </summary>
        [Output("certificationAuthority")]
        public Output<string> CertificationAuthority { get; private set; } = null!;

        /// <summary>
        /// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
        /// </summary>
        [Output("clusterDeployType")]
        public Output<string> ClusterDeployType { get; private set; } = null!;

        /// <summary>
        /// External network address to access.
        /// </summary>
        [Output("clusterExternalEndpoint")]
        public Output<string> ClusterExternalEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specify cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Open internet access or not.
        /// </summary>
        [Output("clusterInternet")]
        public Output<bool?> ClusterInternet { get; private set; } = null!;

        /// <summary>
        /// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the cluster_external_endpoint value may be changed automatically too.
        /// </summary>
        [Output("clusterInternetDomain")]
        public Output<string?> ClusterInternetDomain { get; private set; } = null!;

        /// <summary>
        /// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
        /// </summary>
        [Output("clusterInternetSecurityGroup")]
        public Output<string?> ClusterInternetSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// Open intranet access or not.
        /// </summary>
        [Output("clusterIntranet")]
        public Output<bool?> ClusterIntranet { get; private set; } = null!;

        /// <summary>
        /// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgw_endpoint value may be changed automatically too.
        /// </summary>
        [Output("clusterIntranetDomain")]
        public Output<string?> ClusterIntranetDomain { get; private set; } = null!;

        /// <summary>
        /// Subnet id who can access this independent cluster, this field must and can only set  when `cluster_intranet` is true. `cluster_intranet_subnet_id` can not modify once be set.
        /// </summary>
        [Output("clusterIntranetSubnetId")]
        public Output<string?> ClusterIntranetSubnetId { get; private set; } = null!;

        /// <summary>
        /// Domain name for access.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// this argument was deprecated, use `cluster_internet_security_group` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `cluster_deploy_type` is 'MANAGED_CLUSTER' and `cluster_internet` is true. `managed_cluster_internet_security_policies` can not delete or empty once be set.
        /// </summary>
        [Output("managedClusterInternetSecurityPolicies")]
        public Output<ImmutableArray<string>> ManagedClusterInternetSecurityPolicies { get; private set; } = null!;

        /// <summary>
        /// Password of account.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The Intranet address used for access.
        /// </summary>
        [Output("pgwEndpoint")]
        public Output<string> PgwEndpoint { get; private set; } = null!;

        /// <summary>
        /// User name of account.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a ClusterEndpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterEndpoint(string name, ClusterEndpointArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint", name, args ?? new ClusterEndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClusterEndpoint(string name, Input<string> id, ClusterEndpointState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/clusterEndpoint:ClusterEndpoint", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClusterEndpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClusterEndpoint Get(string name, Input<string> id, ClusterEndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new ClusterEndpoint(name, id, state, options);
        }
    }

    public sealed class ClusterEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Open internet access or not.
        /// </summary>
        [Input("clusterInternet")]
        public Input<bool>? ClusterInternet { get; set; }

        /// <summary>
        /// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the cluster_external_endpoint value may be changed automatically too.
        /// </summary>
        [Input("clusterInternetDomain")]
        public Input<string>? ClusterInternetDomain { get; set; }

        /// <summary>
        /// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
        /// </summary>
        [Input("clusterInternetSecurityGroup")]
        public Input<string>? ClusterInternetSecurityGroup { get; set; }

        /// <summary>
        /// Open intranet access or not.
        /// </summary>
        [Input("clusterIntranet")]
        public Input<bool>? ClusterIntranet { get; set; }

        /// <summary>
        /// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgw_endpoint value may be changed automatically too.
        /// </summary>
        [Input("clusterIntranetDomain")]
        public Input<string>? ClusterIntranetDomain { get; set; }

        /// <summary>
        /// Subnet id who can access this independent cluster, this field must and can only set  when `cluster_intranet` is true. `cluster_intranet_subnet_id` can not modify once be set.
        /// </summary>
        [Input("clusterIntranetSubnetId")]
        public Input<string>? ClusterIntranetSubnetId { get; set; }

        [Input("managedClusterInternetSecurityPolicies")]
        private InputList<string>? _managedClusterInternetSecurityPolicies;

        /// <summary>
        /// this argument was deprecated, use `cluster_internet_security_group` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `cluster_deploy_type` is 'MANAGED_CLUSTER' and `cluster_internet` is true. `managed_cluster_internet_security_policies` can not delete or empty once be set.
        /// </summary>
        [Obsolete(@"this argument was deprecated, use `cluster_internet_security_group` instead.")]
        public InputList<string> ManagedClusterInternetSecurityPolicies
        {
            get => _managedClusterInternetSecurityPolicies ?? (_managedClusterInternetSecurityPolicies = new InputList<string>());
            set => _managedClusterInternetSecurityPolicies = value;
        }

        public ClusterEndpointArgs()
        {
        }
    }

    public sealed class ClusterEndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate used for access.
        /// </summary>
        [Input("certificationAuthority")]
        public Input<string>? CertificationAuthority { get; set; }

        /// <summary>
        /// Cluster deploy type of `MANAGED_CLUSTER` or `INDEPENDENT_CLUSTER`.
        /// </summary>
        [Input("clusterDeployType")]
        public Input<string>? ClusterDeployType { get; set; }

        /// <summary>
        /// External network address to access.
        /// </summary>
        [Input("clusterExternalEndpoint")]
        public Input<string>? ClusterExternalEndpoint { get; set; }

        /// <summary>
        /// Specify cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Open internet access or not.
        /// </summary>
        [Input("clusterInternet")]
        public Input<bool>? ClusterInternet { get; set; }

        /// <summary>
        /// Domain name for cluster Kube-apiserver internet access.  Be careful if you modify value of this parameter, the cluster_external_endpoint value may be changed automatically too.
        /// </summary>
        [Input("clusterInternetDomain")]
        public Input<string>? ClusterInternetDomain { get; set; }

        /// <summary>
        /// Specify security group, NOTE: This argument must not be empty if cluster internet enabled.
        /// </summary>
        [Input("clusterInternetSecurityGroup")]
        public Input<string>? ClusterInternetSecurityGroup { get; set; }

        /// <summary>
        /// Open intranet access or not.
        /// </summary>
        [Input("clusterIntranet")]
        public Input<bool>? ClusterIntranet { get; set; }

        /// <summary>
        /// Domain name for cluster Kube-apiserver intranet access. Be careful if you modify value of this parameter, the pgw_endpoint value may be changed automatically too.
        /// </summary>
        [Input("clusterIntranetDomain")]
        public Input<string>? ClusterIntranetDomain { get; set; }

        /// <summary>
        /// Subnet id who can access this independent cluster, this field must and can only set  when `cluster_intranet` is true. `cluster_intranet_subnet_id` can not modify once be set.
        /// </summary>
        [Input("clusterIntranetSubnetId")]
        public Input<string>? ClusterIntranetSubnetId { get; set; }

        /// <summary>
        /// Domain name for access.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("managedClusterInternetSecurityPolicies")]
        private InputList<string>? _managedClusterInternetSecurityPolicies;

        /// <summary>
        /// this argument was deprecated, use `cluster_internet_security_group` instead. Security policies for managed cluster internet, like:'192.168.1.0/24' or '113.116.51.27', '0.0.0.0/0' means all. This field can only set when field `cluster_deploy_type` is 'MANAGED_CLUSTER' and `cluster_internet` is true. `managed_cluster_internet_security_policies` can not delete or empty once be set.
        /// </summary>
        [Obsolete(@"this argument was deprecated, use `cluster_internet_security_group` instead.")]
        public InputList<string> ManagedClusterInternetSecurityPolicies
        {
            get => _managedClusterInternetSecurityPolicies ?? (_managedClusterInternetSecurityPolicies = new InputList<string>());
            set => _managedClusterInternetSecurityPolicies = value;
        }

        /// <summary>
        /// Password of account.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The Intranet address used for access.
        /// </summary>
        [Input("pgwEndpoint")]
        public Input<string>? PgwEndpoint { get; set; }

        /// <summary>
        /// User name of account.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ClusterEndpointState()
        {
        }
    }
}
