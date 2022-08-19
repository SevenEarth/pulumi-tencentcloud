// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Scf.Inputs
{

    public sealed class FunctionTriggerGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region of cos bucket. if `type` is `cos`, `cos_region` is required.
        /// </summary>
        [Input("cosRegion")]
        public Input<string>? CosRegion { get; set; }

        /// <summary>
        /// Name of the SCF function trigger, if `type` is `ckafka`, the format of name must be `&lt;ckafkaInstanceId&gt;-&lt;topicId&gt;`; if `type` is `cos`, the name is cos bucket id, other In any case, it can be combined arbitrarily. It can only contain English letters, numbers, connectors and underscores. The maximum length is 100.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// TriggerDesc of the SCF function trigger, parameter format of `timer` is linux cron expression; parameter of `cos` type is json string `{"bucketUrl":"&lt;name-appid&gt;.cos.&lt;region&gt;.myqcloud.com","event":"cos:ObjectCreated:*","filter":{"Prefix":"","Suffix":""}}`, where `bucketUrl` is cos bucket (optional), `event` is the cos event trigger, `Prefix` is the corresponding file prefix filter condition, `Suffix` is the suffix filter condition, if not need filter condition can not pass; `cmq` type does not pass this parameter; `ckafka` type parameter format is json string `{"maxMsgNum":"1","offset":"latest"}`; `apigw` type parameter format is json string `{"api":{"authRequired":"FALSE","requestConfig":{"method":"ANY"},"isIntegratedResponse":"FALSE"},"service":{"serviceId":"service-dqzh68sg"},"release":{"environmentName":"test"}}`.
        /// </summary>
        [Input("triggerDesc", required: true)]
        public Input<string> TriggerDesc { get; set; } = null!;

        /// <summary>
        /// Type of the SCF function trigger, support `cos`, `cmq`, `timer`, `ckafka`, `apigw`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionTriggerGetArgs()
        {
        }
    }
}
