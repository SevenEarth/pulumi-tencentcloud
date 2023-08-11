# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'EventTargetTargetDescriptionArgs',
    'EventTargetTargetDescriptionCkafkaTargetParamsArgs',
    'EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs',
    'EventTargetTargetDescriptionEsTargetParamsArgs',
    'EventTargetTargetDescriptionScfParamsArgs',
    'EventTransformTransformationArgs',
    'EventTransformTransformationEtlFilterArgs',
    'EventTransformTransformationExtractionArgs',
    'EventTransformTransformationExtractionTextParamsArgs',
    'EventTransformTransformationTransformArgs',
    'EventTransformTransformationTransformOutputStructArgs',
    'PutEventsEventListArgs',
    'GetBusFilterArgs',
    'GetSearchFilterArgs',
    'GetSearchFilterFilterArgs',
]

@pulumi.input_type
class EventTargetTargetDescriptionArgs:
    def __init__(__self__, *,
                 resource_description: pulumi.Input[str],
                 ckafka_target_params: Optional[pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsArgs']] = None,
                 es_target_params: Optional[pulumi.Input['EventTargetTargetDescriptionEsTargetParamsArgs']] = None,
                 scf_params: Optional[pulumi.Input['EventTargetTargetDescriptionScfParamsArgs']] = None):
        """
        :param pulumi.Input[str] resource_description: QCS resource six-stage format, more references [resource six-stage format](https://cloud.tencent.com/document/product/598/10606).
        :param pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsArgs'] ckafka_target_params: Ckafka parameters.
        :param pulumi.Input['EventTargetTargetDescriptionEsTargetParamsArgs'] es_target_params: ElasticSearch parameters.
        :param pulumi.Input['EventTargetTargetDescriptionScfParamsArgs'] scf_params: cloud function parameters.
        """
        pulumi.set(__self__, "resource_description", resource_description)
        if ckafka_target_params is not None:
            pulumi.set(__self__, "ckafka_target_params", ckafka_target_params)
        if es_target_params is not None:
            pulumi.set(__self__, "es_target_params", es_target_params)
        if scf_params is not None:
            pulumi.set(__self__, "scf_params", scf_params)

    @property
    @pulumi.getter(name="resourceDescription")
    def resource_description(self) -> pulumi.Input[str]:
        """
        QCS resource six-stage format, more references [resource six-stage format](https://cloud.tencent.com/document/product/598/10606).
        """
        return pulumi.get(self, "resource_description")

    @resource_description.setter
    def resource_description(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_description", value)

    @property
    @pulumi.getter(name="ckafkaTargetParams")
    def ckafka_target_params(self) -> Optional[pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsArgs']]:
        """
        Ckafka parameters.
        """
        return pulumi.get(self, "ckafka_target_params")

    @ckafka_target_params.setter
    def ckafka_target_params(self, value: Optional[pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsArgs']]):
        pulumi.set(self, "ckafka_target_params", value)

    @property
    @pulumi.getter(name="esTargetParams")
    def es_target_params(self) -> Optional[pulumi.Input['EventTargetTargetDescriptionEsTargetParamsArgs']]:
        """
        ElasticSearch parameters.
        """
        return pulumi.get(self, "es_target_params")

    @es_target_params.setter
    def es_target_params(self, value: Optional[pulumi.Input['EventTargetTargetDescriptionEsTargetParamsArgs']]):
        pulumi.set(self, "es_target_params", value)

    @property
    @pulumi.getter(name="scfParams")
    def scf_params(self) -> Optional[pulumi.Input['EventTargetTargetDescriptionScfParamsArgs']]:
        """
        cloud function parameters.
        """
        return pulumi.get(self, "scf_params")

    @scf_params.setter
    def scf_params(self, value: Optional[pulumi.Input['EventTargetTargetDescriptionScfParamsArgs']]):
        pulumi.set(self, "scf_params", value)


@pulumi.input_type
class EventTargetTargetDescriptionCkafkaTargetParamsArgs:
    def __init__(__self__, *,
                 retry_policy: pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs'],
                 topic_name: pulumi.Input[str]):
        """
        :param pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs'] retry_policy: retry strategy.
        :param pulumi.Input[str] topic_name: The ckafka topic to deliver to.
        """
        pulumi.set(__self__, "retry_policy", retry_policy)
        pulumi.set(__self__, "topic_name", topic_name)

    @property
    @pulumi.getter(name="retryPolicy")
    def retry_policy(self) -> pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs']:
        """
        retry strategy.
        """
        return pulumi.get(self, "retry_policy")

    @retry_policy.setter
    def retry_policy(self, value: pulumi.Input['EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs']):
        pulumi.set(self, "retry_policy", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The ckafka topic to deliver to.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)


@pulumi.input_type
class EventTargetTargetDescriptionCkafkaTargetParamsRetryPolicyArgs:
    def __init__(__self__, *,
                 max_retry_attempts: pulumi.Input[int],
                 retry_interval: pulumi.Input[int]):
        """
        :param pulumi.Input[int] max_retry_attempts: Maximum number of retries.
        :param pulumi.Input[int] retry_interval: Retry Interval Unit: Seconds.
        """
        pulumi.set(__self__, "max_retry_attempts", max_retry_attempts)
        pulumi.set(__self__, "retry_interval", retry_interval)

    @property
    @pulumi.getter(name="maxRetryAttempts")
    def max_retry_attempts(self) -> pulumi.Input[int]:
        """
        Maximum number of retries.
        """
        return pulumi.get(self, "max_retry_attempts")

    @max_retry_attempts.setter
    def max_retry_attempts(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_retry_attempts", value)

    @property
    @pulumi.getter(name="retryInterval")
    def retry_interval(self) -> pulumi.Input[int]:
        """
        Retry Interval Unit: Seconds.
        """
        return pulumi.get(self, "retry_interval")

    @retry_interval.setter
    def retry_interval(self, value: pulumi.Input[int]):
        pulumi.set(self, "retry_interval", value)


@pulumi.input_type
class EventTargetTargetDescriptionEsTargetParamsArgs:
    def __init__(__self__, *,
                 index_prefix: pulumi.Input[str],
                 index_suffix_mode: pulumi.Input[str],
                 net_mode: pulumi.Input[str],
                 output_mode: pulumi.Input[str],
                 rotation_interval: pulumi.Input[str],
                 index_template_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] index_prefix: index prefix.
        :param pulumi.Input[str] index_suffix_mode: DTS index configuration.
        :param pulumi.Input[str] net_mode: network connection type.
        :param pulumi.Input[str] output_mode: DTS event configuration.
        :param pulumi.Input[str] rotation_interval: es log rotation granularity.
        :param pulumi.Input[str] index_template_type: es template type.
        """
        pulumi.set(__self__, "index_prefix", index_prefix)
        pulumi.set(__self__, "index_suffix_mode", index_suffix_mode)
        pulumi.set(__self__, "net_mode", net_mode)
        pulumi.set(__self__, "output_mode", output_mode)
        pulumi.set(__self__, "rotation_interval", rotation_interval)
        if index_template_type is not None:
            pulumi.set(__self__, "index_template_type", index_template_type)

    @property
    @pulumi.getter(name="indexPrefix")
    def index_prefix(self) -> pulumi.Input[str]:
        """
        index prefix.
        """
        return pulumi.get(self, "index_prefix")

    @index_prefix.setter
    def index_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_prefix", value)

    @property
    @pulumi.getter(name="indexSuffixMode")
    def index_suffix_mode(self) -> pulumi.Input[str]:
        """
        DTS index configuration.
        """
        return pulumi.get(self, "index_suffix_mode")

    @index_suffix_mode.setter
    def index_suffix_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_suffix_mode", value)

    @property
    @pulumi.getter(name="netMode")
    def net_mode(self) -> pulumi.Input[str]:
        """
        network connection type.
        """
        return pulumi.get(self, "net_mode")

    @net_mode.setter
    def net_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "net_mode", value)

    @property
    @pulumi.getter(name="outputMode")
    def output_mode(self) -> pulumi.Input[str]:
        """
        DTS event configuration.
        """
        return pulumi.get(self, "output_mode")

    @output_mode.setter
    def output_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "output_mode", value)

    @property
    @pulumi.getter(name="rotationInterval")
    def rotation_interval(self) -> pulumi.Input[str]:
        """
        es log rotation granularity.
        """
        return pulumi.get(self, "rotation_interval")

    @rotation_interval.setter
    def rotation_interval(self, value: pulumi.Input[str]):
        pulumi.set(self, "rotation_interval", value)

    @property
    @pulumi.getter(name="indexTemplateType")
    def index_template_type(self) -> Optional[pulumi.Input[str]]:
        """
        es template type.
        """
        return pulumi.get(self, "index_template_type")

    @index_template_type.setter
    def index_template_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "index_template_type", value)


@pulumi.input_type
class EventTargetTargetDescriptionScfParamsArgs:
    def __init__(__self__, *,
                 batch_event_count: Optional[pulumi.Input[int]] = None,
                 batch_timeout: Optional[pulumi.Input[int]] = None,
                 enable_batch_delivery: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[int] batch_event_count: Maximum number of events for batch delivery.
        :param pulumi.Input[int] batch_timeout: Maximum waiting time for bulk delivery.
        :param pulumi.Input[bool] enable_batch_delivery: Enable batch delivery.
        """
        if batch_event_count is not None:
            pulumi.set(__self__, "batch_event_count", batch_event_count)
        if batch_timeout is not None:
            pulumi.set(__self__, "batch_timeout", batch_timeout)
        if enable_batch_delivery is not None:
            pulumi.set(__self__, "enable_batch_delivery", enable_batch_delivery)

    @property
    @pulumi.getter(name="batchEventCount")
    def batch_event_count(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of events for batch delivery.
        """
        return pulumi.get(self, "batch_event_count")

    @batch_event_count.setter
    def batch_event_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "batch_event_count", value)

    @property
    @pulumi.getter(name="batchTimeout")
    def batch_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum waiting time for bulk delivery.
        """
        return pulumi.get(self, "batch_timeout")

    @batch_timeout.setter
    def batch_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "batch_timeout", value)

    @property
    @pulumi.getter(name="enableBatchDelivery")
    def enable_batch_delivery(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable batch delivery.
        """
        return pulumi.get(self, "enable_batch_delivery")

    @enable_batch_delivery.setter
    def enable_batch_delivery(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_batch_delivery", value)


@pulumi.input_type
class EventTransformTransformationArgs:
    def __init__(__self__, *,
                 etl_filter: Optional[pulumi.Input['EventTransformTransformationEtlFilterArgs']] = None,
                 extraction: Optional[pulumi.Input['EventTransformTransformationExtractionArgs']] = None,
                 transform: Optional[pulumi.Input['EventTransformTransformationTransformArgs']] = None):
        """
        :param pulumi.Input['EventTransformTransformationEtlFilterArgs'] etl_filter: Describe how to filter data.
        :param pulumi.Input['EventTransformTransformationExtractionArgs'] extraction: Describe how to extract data.
        :param pulumi.Input['EventTransformTransformationTransformArgs'] transform: Describe how to convert data.
        """
        if etl_filter is not None:
            pulumi.set(__self__, "etl_filter", etl_filter)
        if extraction is not None:
            pulumi.set(__self__, "extraction", extraction)
        if transform is not None:
            pulumi.set(__self__, "transform", transform)

    @property
    @pulumi.getter(name="etlFilter")
    def etl_filter(self) -> Optional[pulumi.Input['EventTransformTransformationEtlFilterArgs']]:
        """
        Describe how to filter data.
        """
        return pulumi.get(self, "etl_filter")

    @etl_filter.setter
    def etl_filter(self, value: Optional[pulumi.Input['EventTransformTransformationEtlFilterArgs']]):
        pulumi.set(self, "etl_filter", value)

    @property
    @pulumi.getter
    def extraction(self) -> Optional[pulumi.Input['EventTransformTransformationExtractionArgs']]:
        """
        Describe how to extract data.
        """
        return pulumi.get(self, "extraction")

    @extraction.setter
    def extraction(self, value: Optional[pulumi.Input['EventTransformTransformationExtractionArgs']]):
        pulumi.set(self, "extraction", value)

    @property
    @pulumi.getter
    def transform(self) -> Optional[pulumi.Input['EventTransformTransformationTransformArgs']]:
        """
        Describe how to convert data.
        """
        return pulumi.get(self, "transform")

    @transform.setter
    def transform(self, value: Optional[pulumi.Input['EventTransformTransformationTransformArgs']]):
        pulumi.set(self, "transform", value)


@pulumi.input_type
class EventTransformTransformationEtlFilterArgs:
    def __init__(__self__, *,
                 filter: pulumi.Input[str]):
        """
        :param pulumi.Input[str] filter: Grammatical Rules are consistent.
        """
        pulumi.set(__self__, "filter", filter)

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Input[str]:
        """
        Grammatical Rules are consistent.
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: pulumi.Input[str]):
        pulumi.set(self, "filter", value)


@pulumi.input_type
class EventTransformTransformationExtractionArgs:
    def __init__(__self__, *,
                 extraction_input_path: pulumi.Input[str],
                 format: pulumi.Input[str],
                 text_params: Optional[pulumi.Input['EventTransformTransformationExtractionTextParamsArgs']] = None):
        """
        :param pulumi.Input[str] extraction_input_path: JsonPath, if not specified, the default value $.
        :param pulumi.Input[str] format: Value: `TEXT`, `JSON`.
        :param pulumi.Input['EventTransformTransformationExtractionTextParamsArgs'] text_params: Only Text needs to be passed.
        """
        pulumi.set(__self__, "extraction_input_path", extraction_input_path)
        pulumi.set(__self__, "format", format)
        if text_params is not None:
            pulumi.set(__self__, "text_params", text_params)

    @property
    @pulumi.getter(name="extractionInputPath")
    def extraction_input_path(self) -> pulumi.Input[str]:
        """
        JsonPath, if not specified, the default value $.
        """
        return pulumi.get(self, "extraction_input_path")

    @extraction_input_path.setter
    def extraction_input_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "extraction_input_path", value)

    @property
    @pulumi.getter
    def format(self) -> pulumi.Input[str]:
        """
        Value: `TEXT`, `JSON`.
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: pulumi.Input[str]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="textParams")
    def text_params(self) -> Optional[pulumi.Input['EventTransformTransformationExtractionTextParamsArgs']]:
        """
        Only Text needs to be passed.
        """
        return pulumi.get(self, "text_params")

    @text_params.setter
    def text_params(self, value: Optional[pulumi.Input['EventTransformTransformationExtractionTextParamsArgs']]):
        pulumi.set(self, "text_params", value)


@pulumi.input_type
class EventTransformTransformationExtractionTextParamsArgs:
    def __init__(__self__, *,
                 regex: Optional[pulumi.Input[str]] = None,
                 separator: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] regex: Fill in the regular expression: length 128.
        :param pulumi.Input[str] separator: `Comma`, `|`, `tab`, `space`, `newline`, `%`, `#`, the limit length is 1.
        """
        if regex is not None:
            pulumi.set(__self__, "regex", regex)
        if separator is not None:
            pulumi.set(__self__, "separator", separator)

    @property
    @pulumi.getter
    def regex(self) -> Optional[pulumi.Input[str]]:
        """
        Fill in the regular expression: length 128.
        """
        return pulumi.get(self, "regex")

    @regex.setter
    def regex(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "regex", value)

    @property
    @pulumi.getter
    def separator(self) -> Optional[pulumi.Input[str]]:
        """
        `Comma`, `|`, `tab`, `space`, `newline`, `%`, `#`, the limit length is 1.
        """
        return pulumi.get(self, "separator")

    @separator.setter
    def separator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "separator", value)


@pulumi.input_type
class EventTransformTransformationTransformArgs:
    def __init__(__self__, *,
                 output_structs: pulumi.Input[Sequence[pulumi.Input['EventTransformTransformationTransformOutputStructArgs']]]):
        """
        :param pulumi.Input[Sequence[pulumi.Input['EventTransformTransformationTransformOutputStructArgs']]] output_structs: Describe how the data is transformed.
        """
        pulumi.set(__self__, "output_structs", output_structs)

    @property
    @pulumi.getter(name="outputStructs")
    def output_structs(self) -> pulumi.Input[Sequence[pulumi.Input['EventTransformTransformationTransformOutputStructArgs']]]:
        """
        Describe how the data is transformed.
        """
        return pulumi.get(self, "output_structs")

    @output_structs.setter
    def output_structs(self, value: pulumi.Input[Sequence[pulumi.Input['EventTransformTransformationTransformOutputStructArgs']]]):
        pulumi.set(self, "output_structs", value)


@pulumi.input_type
class EventTransformTransformationTransformOutputStructArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str],
                 value_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: Corresponding to the key in the output json.
        :param pulumi.Input[str] value: You can fill in the json-path and also support constants or built-in keyword date types.
        :param pulumi.Input[str] value_type: The data type of value, optional values: `STRING`, `NUMBER`, `BOOLEAN`, `NULL`, `SYS_VARIABLE`, `JSONPATH`.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)
        pulumi.set(__self__, "value_type", value_type)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        Corresponding to the key in the output json.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        You can fill in the json-path and also support constants or built-in keyword date types.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> pulumi.Input[str]:
        """
        The data type of value, optional values: `STRING`, `NUMBER`, `BOOLEAN`, `NULL`, `SYS_VARIABLE`, `JSONPATH`.
        """
        return pulumi.get(self, "value_type")

    @value_type.setter
    def value_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "value_type", value)


@pulumi.input_type
class PutEventsEventListArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str],
                 source: pulumi.Input[str],
                 subject: pulumi.Input[str],
                 type: pulumi.Input[str],
                 time: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] data: Event data, the content is controlled by the system that created the event, the current datacontenttype only supports application/json;charset=utf-8, so this field is a json string.
        :param pulumi.Input[str] source: Event source information, new product reporting must comply with EB specifications.
        :param pulumi.Input[str] subject: Detailed description of the event source, customizable, optional. The cloud service defaults to the standard qcs resource representation syntax: qcs::dts:ap-guangzhou:appid/uin:xxx.
        :param pulumi.Input[str] type: Event type, customizable, optional. The cloud service writes COS:Created:PostObject by default, use: to separate the type field.
        :param pulumi.Input[int] time: The timestamp in milliseconds when the event occurred,time.Now().UnixNano()/1e6.
        """
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "subject", subject)
        pulumi.set(__self__, "type", type)
        if time is not None:
            pulumi.set(__self__, "time", time)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        Event data, the content is controlled by the system that created the event, the current datacontenttype only supports application/json;charset=utf-8, so this field is a json string.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        Event source information, new product reporting must comply with EB specifications.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def subject(self) -> pulumi.Input[str]:
        """
        Detailed description of the event source, customizable, optional. The cloud service defaults to the standard qcs resource representation syntax: qcs::dts:ap-guangzhou:appid/uin:xxx.
        """
        return pulumi.get(self, "subject")

    @subject.setter
    def subject(self, value: pulumi.Input[str]):
        pulumi.set(self, "subject", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Event type, customizable, optional. The cloud service writes COS:Created:PostObject by default, use: to separate the type field.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def time(self) -> Optional[pulumi.Input[int]]:
        """
        The timestamp in milliseconds when the event occurred,time.Now().UnixNano()/1e6.
        """
        return pulumi.get(self, "time")

    @time.setter
    def time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time", value)


@pulumi.input_type
class GetBusFilterArgs:
    def __init__(__self__, *,
                 name: str,
                 values: Sequence[str]):
        """
        :param str name: The name of the filter key.
        :param Sequence[str] values: One or more filter values.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the filter key.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: str):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        One or more filter values.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GetSearchFilterArgs:
    def __init__(__self__, *,
                 filters: Optional[Sequence['GetSearchFilterFilterArgs']] = None,
                 key: Optional[str] = None,
                 operator: Optional[str] = None,
                 type: Optional[str] = None,
                 value: Optional[str] = None):
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['GetSearchFilterFilterArgs']]:
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[Sequence['GetSearchFilterFilterArgs']]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[str]:
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[str]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetSearchFilterFilterArgs:
    def __init__(__self__, *,
                 key: str,
                 operator: str,
                 value: str):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "operator", operator)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def operator(self) -> str:
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: str):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


