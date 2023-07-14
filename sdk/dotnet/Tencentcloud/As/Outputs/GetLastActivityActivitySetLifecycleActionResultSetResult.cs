// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.As.Outputs
{

    [OutputType]
    public sealed class GetLastActivityActivitySetLifecycleActionResultSetResult
    {
        /// <summary>
        /// ID of the instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Execution task ID. You can query the result by using the DescribeInvocations API of TAT.
        /// </summary>
        public readonly string InvocationId;
        /// <summary>
        /// Result of command invocation, value range: SUCCESSFUL, FAILED, NONE.
        /// </summary>
        public readonly string InvokeCommandResult;
        /// <summary>
        /// Result of the lifecycle hook action, value range: CONTINUE, ABANDON.
        /// </summary>
        public readonly string LifecycleActionResult;
        /// <summary>
        /// ID of the lifecycle hook.
        /// </summary>
        public readonly string LifecycleHookId;
        /// <summary>
        /// Notification result, which indicates whether it is successful to notify CMQ/TDMQ, value range: SUCCESSFUL, FAILED, NONE.
        /// </summary>
        public readonly string NotificationResult;
        /// <summary>
        /// Reason of the result, value range: HEARTBEAT_TIMEOUT: Heartbeat timed out. The setting of DefaultResult is used. NOTIFICATION_FAILURE: Failed to send the notification. The setting of DefaultResult is used. CALL_INTERFACE: Calls the CompleteLifecycleAction to set the result ANOTHER_ACTION_ABANDON: It has been set to ABANDON by another operation. COMMAND_CALL_FAILURE: Failed to call the command. The DefaultResult is applied. COMMAND_EXEC_FINISH: Command completed COMMAND_CALL_FAILURE: Failed to execute the command. The DefaultResult is applied. COMMAND_EXEC_RESULT_CHECK_FAILURE: Failed to check the command result. The DefaultResult is applied.
        /// </summary>
        public readonly string ResultReason;

        [OutputConstructor]
        private GetLastActivityActivitySetLifecycleActionResultSetResult(
            string instanceId,

            string invocationId,

            string invokeCommandResult,

            string lifecycleActionResult,

            string lifecycleHookId,

            string notificationResult,

            string resultReason)
        {
            InstanceId = instanceId;
            InvocationId = invocationId;
            InvokeCommandResult = invokeCommandResult;
            LifecycleActionResult = lifecycleActionResult;
            LifecycleHookId = lifecycleHookId;
            NotificationResult = notificationResult;
            ResultReason = resultReason;
        }
    }
}
