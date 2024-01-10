// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Antiddos
{
    /// <summary>
    /// Provides a resource to create a antiddos cc_precision_policy
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ccPrecisionPolicy = new Tencentcloud.Antiddos.CcPrecisionPolicy("ccPrecisionPolicy", new Tencentcloud.Antiddos.CcPrecisionPolicyArgs
    ///         {
    ///             Domain = "t.baidu.com",
    ///             InstanceId = "bgpip-0000078h",
    ///             Ip = "212.64.62.191",
    ///             PolicyAction = "drop",
    ///             PolicyLists = 
    ///             {
    ///                 new Tencentcloud.Antiddos.Inputs.CcPrecisionPolicyPolicyListArgs
    ///                 {
    ///                     FieldName = "cgi",
    ///                     FieldType = "value",
    ///                     Value = "a.com",
    ///                     ValueOperator = "equal",
    ///                 },
    ///                 new Tencentcloud.Antiddos.Inputs.CcPrecisionPolicyPolicyListArgs
    ///                 {
    ///                     FieldName = "ua",
    ///                     FieldType = "value",
    ///                     Value = "test",
    ///                     ValueOperator = "equal",
    ///                 },
    ///             },
    ///             Protocol = "http",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// antiddos cc_precision_policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Antiddos/ccPrecisionPolicy:CcPrecisionPolicy cc_precision_policy ${instanceId}#${policyId}#${instanceIp}#${domain}#${protocol}
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Antiddos/ccPrecisionPolicy:CcPrecisionPolicy")]
    public partial class CcPrecisionPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Ip value.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// policy type, alg or drop.
        /// </summary>
        [Output("policyAction")]
        public Output<string> PolicyAction { get; private set; } = null!;

        /// <summary>
        /// policy list.
        /// </summary>
        [Output("policyLists")]
        public Output<ImmutableArray<Outputs.CcPrecisionPolicyPolicyList>> PolicyLists { get; private set; } = null!;

        /// <summary>
        /// protocol http or https.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;


        /// <summary>
        /// Create a CcPrecisionPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CcPrecisionPolicy(string name, CcPrecisionPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ccPrecisionPolicy:CcPrecisionPolicy", name, args ?? new CcPrecisionPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CcPrecisionPolicy(string name, Input<string> id, CcPrecisionPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Antiddos/ccPrecisionPolicy:CcPrecisionPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CcPrecisionPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CcPrecisionPolicy Get(string name, Input<string> id, CcPrecisionPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new CcPrecisionPolicy(name, id, state, options);
        }
    }

    public sealed class CcPrecisionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Ip value.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// policy type, alg or drop.
        /// </summary>
        [Input("policyAction", required: true)]
        public Input<string> PolicyAction { get; set; } = null!;

        [Input("policyLists", required: true)]
        private InputList<Inputs.CcPrecisionPolicyPolicyListArgs>? _policyLists;

        /// <summary>
        /// policy list.
        /// </summary>
        public InputList<Inputs.CcPrecisionPolicyPolicyListArgs> PolicyLists
        {
            get => _policyLists ?? (_policyLists = new InputList<Inputs.CcPrecisionPolicyPolicyListArgs>());
            set => _policyLists = value;
        }

        /// <summary>
        /// protocol http or https.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public CcPrecisionPolicyArgs()
        {
        }
    }

    public sealed class CcPrecisionPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// domain.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Ip value.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// policy type, alg or drop.
        /// </summary>
        [Input("policyAction")]
        public Input<string>? PolicyAction { get; set; }

        [Input("policyLists")]
        private InputList<Inputs.CcPrecisionPolicyPolicyListGetArgs>? _policyLists;

        /// <summary>
        /// policy list.
        /// </summary>
        public InputList<Inputs.CcPrecisionPolicyPolicyListGetArgs> PolicyLists
        {
            get => _policyLists ?? (_policyLists = new InputList<Inputs.CcPrecisionPolicyPolicyListGetArgs>());
            set => _policyLists = value;
        }

        /// <summary>
        /// protocol http or https.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public CcPrecisionPolicyState()
        {
        }
    }
}
