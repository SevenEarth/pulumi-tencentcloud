// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.As
{
    /// <summary>
    /// Provides a resource for an AS (Auto scaling) lifecycle hook.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = Pulumi.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var lifecycleHook = new Tencentcloud.As.LifecycleHook("lifecycleHook", new Tencentcloud.As.LifecycleHookArgs
    ///         {
    ///             DefaultResult = "CONTINUE",
    ///             HeartbeatTimeout = 500,
    ///             LifecycleHookName = "tf-as-lifecycle-hook",
    ///             LifecycleTransition = "INSTANCE_LAUNCHING",
    ///             NotificationMetadata = "tf test",
    ///             NotificationQueueName = "lifcyclehook",
    ///             NotificationTargetType = "CMQ_QUEUE",
    ///             ScalingGroupId = "sg-12af45",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:As/lifecycleHook:LifecycleHook")]
    public partial class LifecycleHook : Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        /// </summary>
        [Output("defaultResult")]
        public Output<string?> DefaultResult { get; private set; } = null!;

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        /// </summary>
        [Output("heartbeatTimeout")]
        public Output<int?> HeartbeatTimeout { get; private set; } = null!;

        /// <summary>
        /// The name of the lifecycle hook.
        /// </summary>
        [Output("lifecycleHookName")]
        public Output<string> LifecycleHookName { get; private set; } = null!;

        /// <summary>
        /// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        /// </summary>
        [Output("lifecycleTransition")]
        public Output<string> LifecycleTransition { get; private set; } = null!;

        /// <summary>
        /// Contains additional information that you want to include any time AS sends a message to the notification target.
        /// </summary>
        [Output("notificationMetadata")]
        public Output<string?> NotificationMetadata { get; private set; } = null!;

        /// <summary>
        /// For CMQ_QUEUE type, a name of queue must be set.
        /// </summary>
        [Output("notificationQueueName")]
        public Output<string?> NotificationQueueName { get; private set; } = null!;

        /// <summary>
        /// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        /// </summary>
        [Output("notificationTargetType")]
        public Output<string?> NotificationTargetType { get; private set; } = null!;

        /// <summary>
        /// For CMQ_TOPIC type, a name of topic must be set.
        /// </summary>
        [Output("notificationTopicName")]
        public Output<string?> NotificationTopicName { get; private set; } = null!;

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Output("scalingGroupId")]
        public Output<string> ScalingGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a LifecycleHook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LifecycleHook(string name, LifecycleHookArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:As/lifecycleHook:LifecycleHook", name, args ?? new LifecycleHookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LifecycleHook(string name, Input<string> id, LifecycleHookState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:As/lifecycleHook:LifecycleHook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LifecycleHook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LifecycleHook Get(string name, Input<string> id, LifecycleHookState? state = null, CustomResourceOptions? options = null)
        {
            return new LifecycleHook(name, id, state, options);
        }
    }

    public sealed class LifecycleHookArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        /// </summary>
        [Input("defaultResult")]
        public Input<string>? DefaultResult { get; set; }

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        /// </summary>
        [Input("heartbeatTimeout")]
        public Input<int>? HeartbeatTimeout { get; set; }

        /// <summary>
        /// The name of the lifecycle hook.
        /// </summary>
        [Input("lifecycleHookName", required: true)]
        public Input<string> LifecycleHookName { get; set; } = null!;

        /// <summary>
        /// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        /// </summary>
        [Input("lifecycleTransition", required: true)]
        public Input<string> LifecycleTransition { get; set; } = null!;

        /// <summary>
        /// Contains additional information that you want to include any time AS sends a message to the notification target.
        /// </summary>
        [Input("notificationMetadata")]
        public Input<string>? NotificationMetadata { get; set; }

        /// <summary>
        /// For CMQ_QUEUE type, a name of queue must be set.
        /// </summary>
        [Input("notificationQueueName")]
        public Input<string>? NotificationQueueName { get; set; }

        /// <summary>
        /// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        /// </summary>
        [Input("notificationTargetType")]
        public Input<string>? NotificationTargetType { get; set; }

        /// <summary>
        /// For CMQ_TOPIC type, a name of topic must be set.
        /// </summary>
        [Input("notificationTopicName")]
        public Input<string>? NotificationTopicName { get; set; }

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Input("scalingGroupId", required: true)]
        public Input<string> ScalingGroupId { get; set; } = null!;

        public LifecycleHookArgs()
        {
        }
    }

    public sealed class LifecycleHookState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the action the AS group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. Valid values: `CONTINUE` and `ABANDON`. The default value is `CONTINUE`.
        /// </summary>
        [Input("defaultResult")]
        public Input<string>? DefaultResult { get; set; }

        /// <summary>
        /// Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. Valid value ranges: (30~7200). and default value is `300`.
        /// </summary>
        [Input("heartbeatTimeout")]
        public Input<int>? HeartbeatTimeout { get; set; }

        /// <summary>
        /// The name of the lifecycle hook.
        /// </summary>
        [Input("lifecycleHookName")]
        public Input<string>? LifecycleHookName { get; set; }

        /// <summary>
        /// The instance state to which you want to attach the lifecycle hook. Valid values: `INSTANCE_LAUNCHING` and `INSTANCE_TERMINATING`.
        /// </summary>
        [Input("lifecycleTransition")]
        public Input<string>? LifecycleTransition { get; set; }

        /// <summary>
        /// Contains additional information that you want to include any time AS sends a message to the notification target.
        /// </summary>
        [Input("notificationMetadata")]
        public Input<string>? NotificationMetadata { get; set; }

        /// <summary>
        /// For CMQ_QUEUE type, a name of queue must be set.
        /// </summary>
        [Input("notificationQueueName")]
        public Input<string>? NotificationQueueName { get; set; }

        /// <summary>
        /// Target type. Valid values: `CMQ_QUEUE`, `CMQ_TOPIC`.
        /// </summary>
        [Input("notificationTargetType")]
        public Input<string>? NotificationTargetType { get; set; }

        /// <summary>
        /// For CMQ_TOPIC type, a name of topic must be set.
        /// </summary>
        [Input("notificationTopicName")]
        public Input<string>? NotificationTopicName { get; set; }

        /// <summary>
        /// ID of a scaling group.
        /// </summary>
        [Input("scalingGroupId")]
        public Input<string>? ScalingGroupId { get; set; }

        public LifecycleHookState()
        {
        }
    }
}
