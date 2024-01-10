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
    /// Provide a resource to configure kubernetes cluster app addons.
    /// 
    /// &gt; **NOTE**: Avoid to using legacy "1.0.0" version, leave the versions empty so we can fetch the latest while creating.
    /// 
    /// ## Example Usage
    /// ### Install cbs addon by passing values
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var addonCbs = new Tencentcloud.Kubernetes.AddonAttachment("addonCbs", new Tencentcloud.Kubernetes.AddonAttachmentArgs
    ///         {
    ///             ClusterId = "cls-xxxxxxxx",
    ///             Values = 
    ///             {
    ///                 "rootdir=/var/lib/kubelet",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Install tcr addon by passing values
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mytcr = new Tencentcloud.Tcr.Instance("mytcr", new Tencentcloud.Tcr.InstanceArgs
    ///         {
    ///             InstanceType = "basic",
    ///             DeleteBucket = true,
    ///             Tags = 
    ///             {
    ///                 { "test", "test" },
    ///             },
    ///         });
    ///         var tcrId = mytcr.Id;
    ///         var tcrName = mytcr.Name;
    ///         var myNs = new Tencentcloud.Tcr.Namespace("myNs", new Tencentcloud.Tcr.NamespaceArgs
    ///         {
    ///             InstanceId = tcrId,
    ///             IsPublic = true,
    ///             IsAutoScan = true,
    ///             IsPreventVul = true,
    ///             Severity = "medium",
    ///             CveWhitelistItems = 
    ///             {
    ///                 new Tencentcloud.Tcr.Inputs.NamespaceCveWhitelistItemArgs
    ///                 {
    ///                     CveId = "cve-xxxxx",
    ///                 },
    ///             },
    ///         });
    ///         var nsName = myNs.Name;
    ///         var myToken = new Tencentcloud.Tcr.Token("myToken", new Tencentcloud.Tcr.TokenArgs
    ///         {
    ///             InstanceId = tcrId,
    ///             Description = "tcr token",
    ///         });
    ///         var userName = myToken.UserName;
    ///         var token = myToken.Token;
    ///         var myIns = Tencentcloud.Tcr.GetInstances.Invoke(new Tencentcloud.Tcr.GetInstancesInvokeArgs
    ///         {
    ///             InstanceId = tcrId,
    ///         });
    ///         var endPoint = myIns.Apply(myIns =&gt; myIns.InstanceLists?[0]?.InternalEndPoint);
    ///         var addonTcr = new Tencentcloud.Kubernetes.AddonAttachment("addonTcr", new Tencentcloud.Kubernetes.AddonAttachmentArgs
    ///         {
    ///             ClusterId = "cls-xxxxxxxx",
    ///             Version = "1.0.0",
    ///             Values = 
    ///             {
    ///                 tcrId.Apply(tcrId =&gt; $"global.imagePullSecretsCrs[0].name={tcrId}-vpc"),
    ///                 nsName.Apply(nsName =&gt; $"global.imagePullSecretsCrs[0].namespaces={nsName}"),
    ///                 "global.imagePullSecretsCrs[0].serviceAccounts=*",
    ///                 "global.imagePullSecretsCrs[0].type=docker",
    ///                 userName.Apply(userName =&gt; $"global.imagePullSecretsCrs[0].dockerUsername={userName}"),
    ///                 token.Apply(token =&gt; $"global.imagePullSecretsCrs[0].dockerPassword={token}"),
    ///                 tcrName.Apply(tcrName =&gt; $"global.imagePullSecretsCrs[0].dockerServer={tcrName}-vpc.tencentcloudcr.com"),
    ///                 tcrId.Apply(tcrId =&gt; $"global.imagePullSecretsCrs[1].name={tcrId}-public"),
    ///                 nsName.Apply(nsName =&gt; $"global.imagePullSecretsCrs[1].namespaces={nsName}"),
    ///                 "global.imagePullSecretsCrs[1].serviceAccounts=*",
    ///                 "global.imagePullSecretsCrs[1].type=docker",
    ///                 userName.Apply(userName =&gt; $"global.imagePullSecretsCrs[1].dockerUsername={userName}"),
    ///                 token.Apply(token =&gt; $"global.imagePullSecretsCrs[1].dockerPassword={token}"),
    ///                 tcrName.Apply(tcrName =&gt; $"global.imagePullSecretsCrs[1].dockerServer={tcrName}.tencentcloudcr.com"),
    ///                 "global.cluster.region=gz",
    ///                 "global.cluster.longregion=ap-guangzhou",
    ///                 tcrName.Apply(tcrName =&gt; $"global.hosts[0].domain={tcrName}-vpc.tencentcloudcr.com"),
    ///                 endPoint.Apply(endPoint =&gt; $"global.hosts[0].ip={endPoint}"),
    ///                 "global.hosts[0].disabled=false",
    ///                 tcrName.Apply(tcrName =&gt; $"global.hosts[1].domain={tcrName}.tencentcloudcr.com"),
    ///                 endPoint.Apply(endPoint =&gt; $"global.hosts[1].ip={endPoint}"),
    ///                 "global.hosts[1].disabled=false",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Install new addon by passing spec json to req_body directly
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var addonCbs = new Tencentcloud.Kubernetes.AddonAttachment("addonCbs", new Tencentcloud.Kubernetes.AddonAttachmentArgs
    ///         {
    ///             ClusterId = "cls-xxxxxxxx",
    ///             RequestBody = @"  {
    ///     ""spec"":{
    ///         ""chart"":{
    ///             ""chartName"":""cbs"",
    ///             ""chartVersion"":""1.0.5""
    ///         },
    ///         ""values"":{
    ///             ""rawValuesType"":""yaml"",
    ///             ""values"":[
    ///               ""rootdir=/var/lib/kubelet""
    ///             ]
    ///         }
    ///     }
    ///   }
    /// 
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Addon can be imported by using cluster_id#addon_name
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Kubernetes/addonAttachment:AddonAttachment addon_cos cls-xxxxxxxx#cos
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Kubernetes/addonAttachment:AddonAttachment")]
    public partial class AddonAttachment : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Name of addon.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Raw Values. Conflict with `request_body`. Required with `raw_values_type`.
        /// </summary>
        [Output("rawValues")]
        public Output<string> RawValues { get; private set; } = null!;

        /// <summary>
        /// The type of raw Values. Required with `raw_values`.
        /// </summary>
        [Output("rawValuesType")]
        public Output<string> RawValuesType { get; private set; } = null!;

        /// <summary>
        /// Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
        /// </summary>
        [Output("requestBody")]
        public Output<string?> RequestBody { get; private set; } = null!;

        /// <summary>
        /// Addon response body.
        /// </summary>
        [Output("responseBody")]
        public Output<string> ResponseBody { get; private set; } = null!;

        /// <summary>
        /// Addon current status.
        /// </summary>
        [Output("status")]
        public Output<ImmutableDictionary<string, object>> Status { get; private set; } = null!;

        /// <summary>
        /// Values the addon passthroughs. Conflict with `request_body`.
        /// </summary>
        [Output("values")]
        public Output<ImmutableArray<string>> Values { get; private set; } = null!;

        /// <summary>
        /// Addon version, default latest version. Conflict with `request_body`.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AddonAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AddonAttachment(string name, AddonAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/addonAttachment:AddonAttachment", name, args ?? new AddonAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AddonAttachment(string name, Input<string> id, AddonAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Kubernetes/addonAttachment:AddonAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AddonAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AddonAttachment Get(string name, Input<string> id, AddonAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new AddonAttachment(name, id, state, options);
        }
    }

    public sealed class AddonAttachmentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Raw Values. Conflict with `request_body`. Required with `raw_values_type`.
        /// </summary>
        [Input("rawValues")]
        public Input<string>? RawValues { get; set; }

        /// <summary>
        /// The type of raw Values. Required with `raw_values`.
        /// </summary>
        [Input("rawValuesType")]
        public Input<string>? RawValuesType { get; set; }

        /// <summary>
        /// Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
        /// </summary>
        [Input("requestBody")]
        public Input<string>? RequestBody { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Values the addon passthroughs. Conflict with `request_body`.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        /// <summary>
        /// Addon version, default latest version. Conflict with `request_body`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AddonAttachmentArgs()
        {
        }
    }

    public sealed class AddonAttachmentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Raw Values. Conflict with `request_body`. Required with `raw_values_type`.
        /// </summary>
        [Input("rawValues")]
        public Input<string>? RawValues { get; set; }

        /// <summary>
        /// The type of raw Values. Required with `raw_values`.
        /// </summary>
        [Input("rawValuesType")]
        public Input<string>? RawValuesType { get; set; }

        /// <summary>
        /// Serialized json string as request body of addon spec. If set, will ignore `version` and `values`.
        /// </summary>
        [Input("requestBody")]
        public Input<string>? RequestBody { get; set; }

        /// <summary>
        /// Addon response body.
        /// </summary>
        [Input("responseBody")]
        public Input<string>? ResponseBody { get; set; }

        [Input("status")]
        private InputMap<object>? _status;

        /// <summary>
        /// Addon current status.
        /// </summary>
        public InputMap<object> Status
        {
            get => _status ?? (_status = new InputMap<object>());
            set => _status = value;
        }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Values the addon passthroughs. Conflict with `request_body`.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        /// <summary>
        /// Addon version, default latest version. Conflict with `request_body`.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public AddonAttachmentState()
        {
        }
    }
}
