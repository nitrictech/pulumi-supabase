# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DbFunctionArgs', 'DbFunction']

@pulumi.input_type
class DbFunctionArgs:
    def __init__(__self__, *,
                 behaviour: pulumi.Input[str],
                 definition: pulumi.Input[str],
                 language: pulumi.Input[str],
                 project_ref: pulumi.Input[str],
                 return_type: pulumi.Input[str],
                 schema: pulumi.Input[str],
                 security_definer: pulumi.Input[bool],
                 args: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 config_params: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DbFunction resource.
        """
        pulumi.set(__self__, "behaviour", behaviour)
        pulumi.set(__self__, "definition", definition)
        pulumi.set(__self__, "language", language)
        pulumi.set(__self__, "project_ref", project_ref)
        pulumi.set(__self__, "return_type", return_type)
        pulumi.set(__self__, "schema", schema)
        pulumi.set(__self__, "security_definer", security_definer)
        if args is not None:
            pulumi.set(__self__, "args", args)
        if config_params is not None:
            pulumi.set(__self__, "config_params", config_params)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def behaviour(self) -> pulumi.Input[str]:
        return pulumi.get(self, "behaviour")

    @behaviour.setter
    def behaviour(self, value: pulumi.Input[str]):
        pulumi.set(self, "behaviour", value)

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Input[str]:
        return pulumi.get(self, "definition")

    @definition.setter
    def definition(self, value: pulumi.Input[str]):
        pulumi.set(self, "definition", value)

    @property
    @pulumi.getter
    def language(self) -> pulumi.Input[str]:
        return pulumi.get(self, "language")

    @language.setter
    def language(self, value: pulumi.Input[str]):
        pulumi.set(self, "language", value)

    @property
    @pulumi.getter
    def project_ref(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_ref")

    @project_ref.setter
    def project_ref(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_ref", value)

    @property
    @pulumi.getter
    def return_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "return_type")

    @return_type.setter
    def return_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "return_type", value)

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Input[str]:
        return pulumi.get(self, "schema")

    @schema.setter
    def schema(self, value: pulumi.Input[str]):
        pulumi.set(self, "schema", value)

    @property
    @pulumi.getter
    def security_definer(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "security_definer")

    @security_definer.setter
    def security_definer(self, value: pulumi.Input[bool]):
        pulumi.set(self, "security_definer", value)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def config_params(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "config_params")

    @config_params.setter
    def config_params(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config_params", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class DbFunction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 behaviour: Optional[pulumi.Input[str]] = None,
                 config_params: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_ref: Optional[pulumi.Input[str]] = None,
                 return_type: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 security_definer: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a DbFunction resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DbFunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DbFunction resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DbFunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DbFunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 behaviour: Optional[pulumi.Input[str]] = None,
                 config_params: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 definition: Optional[pulumi.Input[str]] = None,
                 language: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_ref: Optional[pulumi.Input[str]] = None,
                 return_type: Optional[pulumi.Input[str]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 security_definer: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DbFunctionArgs.__new__(DbFunctionArgs)

            __props__.__dict__["args"] = args
            if behaviour is None and not opts.urn:
                raise TypeError("Missing required property 'behaviour'")
            __props__.__dict__["behaviour"] = behaviour
            __props__.__dict__["config_params"] = config_params
            if definition is None and not opts.urn:
                raise TypeError("Missing required property 'definition'")
            __props__.__dict__["definition"] = definition
            if language is None and not opts.urn:
                raise TypeError("Missing required property 'language'")
            __props__.__dict__["language"] = language
            __props__.__dict__["name"] = name
            if project_ref is None and not opts.urn:
                raise TypeError("Missing required property 'project_ref'")
            __props__.__dict__["project_ref"] = project_ref
            if return_type is None and not opts.urn:
                raise TypeError("Missing required property 'return_type'")
            __props__.__dict__["return_type"] = return_type
            if schema is None and not opts.urn:
                raise TypeError("Missing required property 'schema'")
            __props__.__dict__["schema"] = schema
            if security_definer is None and not opts.urn:
                raise TypeError("Missing required property 'security_definer'")
            __props__.__dict__["security_definer"] = security_definer
            __props__.__dict__["function_id"] = None
            __props__.__dict__["function_name"] = None
        super(DbFunction, __self__).__init__(
            'supabase:functions:DbFunction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DbFunction':
        """
        Get an existing DbFunction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DbFunctionArgs.__new__(DbFunctionArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["behaviour"] = None
        __props__.__dict__["config_params"] = None
        __props__.__dict__["definition"] = None
        __props__.__dict__["function_id"] = None
        __props__.__dict__["function_name"] = None
        __props__.__dict__["language"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_ref"] = None
        __props__.__dict__["return_type"] = None
        __props__.__dict__["schema"] = None
        __props__.__dict__["security_definer"] = None
        return DbFunction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "args")

    @property
    @pulumi.getter
    def behaviour(self) -> pulumi.Output[str]:
        return pulumi.get(self, "behaviour")

    @property
    @pulumi.getter
    def config_params(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "config_params")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[str]:
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter
    def function_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def function_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter
    def language(self) -> pulumi.Output[str]:
        return pulumi.get(self, "language")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project_ref")

    @property
    @pulumi.getter
    def return_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "return_type")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[str]:
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def security_definer(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "security_definer")

