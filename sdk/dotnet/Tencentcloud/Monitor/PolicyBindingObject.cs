// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource for bind objects to a alarm policy resource.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instances = Tencentcloud.Instances.GetInstance.Invoke();
    /// 
    ///     var policy = new Tencentcloud.Monitor.AlarmPolicy("policy", new()
    ///     {
    ///         PolicyName = "hello",
    ///         MonitorType = "MT_QCE",
    ///         Enable = 1,
    ///         ProjectId = 1244035,
    ///         Namespace = "cvm_device",
    ///         Conditions = new Tencentcloud.Monitor.Inputs.AlarmPolicyConditionsArgs
    ///         {
    ///             IsUnionRule = 1,
    ///             Rules = new[]
    ///             {
    ///                 new Tencentcloud.Monitor.Inputs.AlarmPolicyConditionsRuleArgs
    ///                 {
    ///                     MetricName = "CpuUsage",
    ///                     Period = 60,
    ///                     Operator = "ge",
    ///                     Value = "89.9",
    ///                     ContinuePeriod = 1,
    ///                     NoticeFrequency = 3600,
    ///                     IsPowerNotice = 0,
    ///                 },
    ///             },
    ///         },
    ///         EventConditions = new[]
    ///         {
    ///             new Tencentcloud.Monitor.Inputs.AlarmPolicyEventConditionArgs
    ///             {
    ///                 MetricName = "ping_unreachable",
    ///             },
    ///             new Tencentcloud.Monitor.Inputs.AlarmPolicyEventConditionArgs
    ///             {
    ///                 MetricName = "guest_reboot",
    ///             },
    ///         },
    ///         NoticeIds = new[]
    ///         {
    ///             "notice-l9ziyxw6",
    ///         },
    ///         TriggerTasks = new[]
    ///         {
    ///             new Tencentcloud.Monitor.Inputs.AlarmPolicyTriggerTaskArgs
    ///             {
    ///                 Type = "AS",
    ///                 TaskConfig = "{\"Region\":\"ap-guangzhou\",\"Group\":\"asg-0z312312x\",\"Policy\":\"asp-ganig28\"}",
    ///             },
    ///         },
    ///     });
    /// 
    ///     //for cvm
    ///     var binding = new Tencentcloud.Monitor.PolicyBindingObject("binding", new()
    ///     {
    ///         PolicyId = policy.Id,
    ///         Dimensions = new[]
    ///         {
    ///             new Tencentcloud.Monitor.Inputs.PolicyBindingObjectDimensionArgs
    ///             {
    ///                 DimensionsJson = $"{{\"unInstanceId\":\"{instances.Apply(getInstanceResult =&gt; getInstanceResult.InstanceLists[0]?.InstanceId)}\"}}",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Monitor Policy Binding Object can be imported, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Monitor/policyBindingObject:PolicyBindingObject binding policyId
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject")]
    public partial class PolicyBindingObject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        [Output("dimensions")]
        public Output<ImmutableArray<Outputs.PolicyBindingObjectDimension>> Dimensions { get; private set; } = null!;

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyBindingObject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyBindingObject(string name, PolicyBindingObjectArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, args ?? new PolicyBindingObjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyBindingObject(string name, Input<string> id, PolicyBindingObjectState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/policyBindingObject:PolicyBindingObject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyBindingObject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyBindingObject Get(string name, Input<string> id, PolicyBindingObjectState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyBindingObject(name, id, state, options);
        }
    }

    public sealed class PolicyBindingObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("dimensions", required: true)]
        private InputList<Inputs.PolicyBindingObjectDimensionArgs>? _dimensions;

        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.PolicyBindingObjectDimensionArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.PolicyBindingObjectDimensionArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        public PolicyBindingObjectArgs()
        {
        }
        public static new PolicyBindingObjectArgs Empty => new PolicyBindingObjectArgs();
    }

    public sealed class PolicyBindingObjectState : global::Pulumi.ResourceArgs
    {
        [Input("dimensions")]
        private InputList<Inputs.PolicyBindingObjectDimensionGetArgs>? _dimensions;

        /// <summary>
        /// A list objects. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.PolicyBindingObjectDimensionGetArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.PolicyBindingObjectDimensionGetArgs>());
            set => _dimensions = value;
        }

        /// <summary>
        /// Alarm policy ID for binding objects.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        public PolicyBindingObjectState()
        {
        }
        public static new PolicyBindingObjectState Empty => new PolicyBindingObjectState();
    }
}
