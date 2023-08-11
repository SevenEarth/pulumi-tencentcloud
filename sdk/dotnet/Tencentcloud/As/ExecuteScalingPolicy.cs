// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As
{
    /// <summary>
    /// Provides a resource to create a as execute_scaling_policy
    /// 
    /// ## Example Usage
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
    ///         var zones = Output.Create(Tencentcloud.Availability.GetZonesByProduct.InvokeAsync(new Tencentcloud.Availability.GetZonesByProductArgs
    ///         {
    ///             Product = "as",
    ///         }));
    ///         var image = Output.Create(Tencentcloud.Images.GetInstance.InvokeAsync(new Tencentcloud.Images.GetInstanceArgs
    ///         {
    ///             ImageTypes = 
    ///             {
    ///                 "PUBLIC_IMAGE",
    ///             },
    ///             OsName = "TencentOS Server 3.2 (Final)",
    ///         }));
    ///         var vpc = new Tencentcloud.Vpc.Instance("vpc", new Tencentcloud.Vpc.InstanceArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var subnet = new Tencentcloud.Subnet.Instance("subnet", new Tencentcloud.Subnet.InstanceArgs
    ///         {
    ///             VpcId = vpc.Id,
    ///             CidrBlock = "10.0.0.0/16",
    ///             AvailabilityZone = zones.Apply(zones =&gt; zones.Zones?[0]?.Name),
    ///         });
    ///         var exampleScalingConfig = new Tencentcloud.As.ScalingConfig("exampleScalingConfig", new Tencentcloud.As.ScalingConfigArgs
    ///         {
    ///             ConfigurationName = "tf-example",
    ///             ImageId = image.Apply(image =&gt; image.Images?[0]?.ImageId),
    ///             InstanceTypes = 
    ///             {
    ///                 "SA1.SMALL1",
    ///                 "SA2.SMALL1",
    ///                 "SA2.SMALL2",
    ///                 "SA2.SMALL4",
    ///             },
    ///             InstanceNameSettings = new Tencentcloud.As.Inputs.ScalingConfigInstanceNameSettingsArgs
    ///             {
    ///                 InstanceName = "test-ins-name",
    ///             },
    ///         });
    ///         var exampleScalingGroup = new Tencentcloud.As.ScalingGroup("exampleScalingGroup", new Tencentcloud.As.ScalingGroupArgs
    ///         {
    ///             ScalingGroupName = "tf-example",
    ///             ConfigurationId = exampleScalingConfig.Id,
    ///             MaxSize = 4,
    ///             MinSize = 1,
    ///             DesiredCapacity = 2,
    ///             VpcId = vpc.Id,
    ///             SubnetIds = 
    ///             {
    ///                 subnet.Id,
    ///             },
    ///         });
    ///         var exampleScalingPolicy = new Tencentcloud.As.ScalingPolicy("exampleScalingPolicy", new Tencentcloud.As.ScalingPolicyArgs
    ///         {
    ///             ScalingGroupId = exampleScalingGroup.Id,
    ///             PolicyName = "tf-as-scaling-policy",
    ///             AdjustmentType = "EXACT_CAPACITY",
    ///             AdjustmentValue = 0,
    ///             ComparisonOperator = "GREATER_THAN",
    ///             MetricName = "CPU_UTILIZATION",
    ///             Threshold = 80,
    ///             Period = 300,
    ///             ContinuousTime = 10,
    ///             Statistic = "AVERAGE",
    ///             Cooldown = 360,
    ///         });
    ///         var exampleExecuteScalingPolicy = new Tencentcloud.As.ExecuteScalingPolicy("exampleExecuteScalingPolicy", new Tencentcloud.As.ExecuteScalingPolicyArgs
    ///         {
    ///             AutoScalingPolicyId = exampleScalingPolicy.Id,
    ///             HonorCooldown = false,
    ///             TriggerSource = "API",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// as execute_scaling_policy can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy execute_scaling_policy execute_scaling_policy_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy")]
    public partial class ExecuteScalingPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        /// </summary>
        [Output("autoScalingPolicyId")]
        public Output<string> AutoScalingPolicyId { get; private set; } = null!;

        /// <summary>
        /// Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        /// </summary>
        [Output("honorCooldown")]
        public Output<bool?> HonorCooldown { get; private set; } = null!;

        /// <summary>
        /// Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        /// </summary>
        [Output("triggerSource")]
        public Output<string?> TriggerSource { get; private set; } = null!;


        /// <summary>
        /// Create a ExecuteScalingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExecuteScalingPolicy(string name, ExecuteScalingPolicyArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy", name, args ?? new ExecuteScalingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExecuteScalingPolicy(string name, Input<string> id, ExecuteScalingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:As/executeScalingPolicy:ExecuteScalingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExecuteScalingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExecuteScalingPolicy Get(string name, Input<string> id, ExecuteScalingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ExecuteScalingPolicy(name, id, state, options);
        }
    }

    public sealed class ExecuteScalingPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        /// </summary>
        [Input("autoScalingPolicyId", required: true)]
        public Input<string> AutoScalingPolicyId { get; set; } = null!;

        /// <summary>
        /// Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        /// </summary>
        [Input("honorCooldown")]
        public Input<bool>? HonorCooldown { get; set; }

        /// <summary>
        /// Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        /// </summary>
        [Input("triggerSource")]
        public Input<string>? TriggerSource { get; set; }

        public ExecuteScalingPolicyArgs()
        {
        }
    }

    public sealed class ExecuteScalingPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Auto-scaling policy ID. This parameter is not available to a target tracking policy.
        /// </summary>
        [Input("autoScalingPolicyId")]
        public Input<string>? AutoScalingPolicyId { get; set; }

        /// <summary>
        /// Whether to check if the auto scaling group is in the cooldown period. Default value: false.
        /// </summary>
        [Input("honorCooldown")]
        public Input<bool>? HonorCooldown { get; set; }

        /// <summary>
        /// Source that triggers the scaling policy. Valid values: API and CLOUD_MONITOR. Default value: API. The value CLOUD_MONITOR is specific to the Cloud Monitor service.
        /// </summary>
        [Input("triggerSource")]
        public Input<string>? TriggerSource { get; set; }

        public ExecuteScalingPolicyState()
        {
        }
    }
}
