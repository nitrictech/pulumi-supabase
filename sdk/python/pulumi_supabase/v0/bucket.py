# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BucketArgs', 'Bucket']

@pulumi.input_type
class BucketArgs:
    def __init__(__self__, *,
                 project_ref: pulumi.Input[str],
                 allowed_mime_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file_size_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Bucket resource.
        """
        pulumi.set(__self__, "project_ref", project_ref)
        if allowed_mime_types is not None:
            pulumi.set(__self__, "allowed_mime_types", allowed_mime_types)
        if file_size_limit is not None:
            pulumi.set(__self__, "file_size_limit", file_size_limit)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if public is not None:
            pulumi.set(__self__, "public", public)

    @property
    @pulumi.getter
    def project_ref(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project_ref")

    @project_ref.setter
    def project_ref(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_ref", value)

    @property
    @pulumi.getter
    def allowed_mime_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "allowed_mime_types")

    @allowed_mime_types.setter
    def allowed_mime_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_mime_types", value)

    @property
    @pulumi.getter
    def file_size_limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "file_size_limit")

    @file_size_limit.setter
    def file_size_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "file_size_limit", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def public(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "public")

    @public.setter
    def public(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "public", value)


class Bucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_mime_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file_size_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_ref: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Bucket resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Bucket resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_mime_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 file_size_limit: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_ref: Optional[pulumi.Input[str]] = None,
                 public: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketArgs.__new__(BucketArgs)

            __props__.__dict__["allowed_mime_types"] = allowed_mime_types
            __props__.__dict__["file_size_limit"] = file_size_limit
            __props__.__dict__["name"] = name
            if project_ref is None and not opts.urn:
                raise TypeError("Missing required property 'project_ref'")
            __props__.__dict__["project_ref"] = project_ref
            __props__.__dict__["public"] = public
        super(Bucket, __self__).__init__(
            'supabase:v0:Bucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Bucket':
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BucketArgs.__new__(BucketArgs)

        __props__.__dict__["allowed_mime_types"] = None
        __props__.__dict__["file_size_limit"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project_ref"] = None
        __props__.__dict__["public"] = None
        return Bucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def allowed_mime_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "allowed_mime_types")

    @property
    @pulumi.getter
    def file_size_limit(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "file_size_limit")

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
    def public(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "public")

