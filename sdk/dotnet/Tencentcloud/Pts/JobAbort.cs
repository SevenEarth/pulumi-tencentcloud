// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts
{
    /// <summary>
    /// Provides a resource to create a pts job_abort
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
    ///         var jobAbort = new Tencentcloud.Pts.JobAbort("jobAbort", new Tencentcloud.Pts.JobAbortArgs
    ///         {
    ///             JobId = "job-my644ozi",
    ///             ProjectId = "project-45vw7v82",
    ///             ScenarioId = "scenario-22q19f3k",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Pts/jobAbort:JobAbort")]
    public partial class JobAbort : Pulumi.CustomResource
    {
        /// <summary>
        /// The reason for aborting the job.
        /// </summary>
        [Output("abortReason")]
        public Output<int?> AbortReason { get; private set; } = null!;

        /// <summary>
        /// Job ID.
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Output("projectId")]
        public Output<string> ProjectId { get; private set; } = null!;

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Output("scenarioId")]
        public Output<string> ScenarioId { get; private set; } = null!;


        /// <summary>
        /// Create a JobAbort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobAbort(string name, JobAbortArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/jobAbort:JobAbort", name, args ?? new JobAbortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobAbort(string name, Input<string> id, JobAbortState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Pts/jobAbort:JobAbort", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing JobAbort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobAbort Get(string name, Input<string> id, JobAbortState? state = null, CustomResourceOptions? options = null)
        {
            return new JobAbort(name, id, state, options);
        }
    }

    public sealed class JobAbortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reason for aborting the job.
        /// </summary>
        [Input("abortReason")]
        public Input<int>? AbortReason { get; set; }

        /// <summary>
        /// Job ID.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId", required: true)]
        public Input<string> ProjectId { get; set; } = null!;

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Input("scenarioId", required: true)]
        public Input<string> ScenarioId { get; set; } = null!;

        public JobAbortArgs()
        {
        }
    }

    public sealed class JobAbortState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reason for aborting the job.
        /// </summary>
        [Input("abortReason")]
        public Input<int>? AbortReason { get; set; }

        /// <summary>
        /// Job ID.
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// Project ID.
        /// </summary>
        [Input("projectId")]
        public Input<string>? ProjectId { get; set; }

        /// <summary>
        /// Scenario ID.
        /// </summary>
        [Input("scenarioId")]
        public Input<string>? ScenarioId { get; set; }

        public JobAbortState()
        {
        }
    }
}