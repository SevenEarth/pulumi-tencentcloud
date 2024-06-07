// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("headSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs>? _headSets;

        /// <summary>
        /// Opening credits list.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs> HeadSets
        {
            get => _headSets ?? (_headSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterHeadSetGetArgs>());
            set => _headSets = value;
        }

        [Input("tailSets")]
        private InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs>? _tailSets;

        /// <summary>
        /// Closing credits list.
        /// </summary>
        public InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs> TailSets
        {
            get => _tailSets ?? (_tailSets = new InputList<Inputs.ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterTailSetGetArgs>());
            set => _tailSets = value;
        }

        public ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs()
        {
        }
        public static new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs Empty => new ProcessMediaOperationMediaProcessTaskTranscodeTaskSetHeadTailParameterGetArgs();
    }
}
