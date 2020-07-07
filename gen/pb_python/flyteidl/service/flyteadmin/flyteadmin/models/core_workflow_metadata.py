# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.core_quality_of_service import CoreQualityOfService  # noqa: F401,E501
from flyteadmin.models.workflow_metadata_on_failure_policy import WorkflowMetadataOnFailurePolicy  # noqa: F401,E501


class CoreWorkflowMetadata(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'queuing_budget': 'CoreQualityOfService',
        'on_failure': 'WorkflowMetadataOnFailurePolicy'
    }

    attribute_map = {
        'queuing_budget': 'queuing_budget',
        'on_failure': 'on_failure'
    }

    def __init__(self, queuing_budget=None, on_failure=None):  # noqa: E501
        """CoreWorkflowMetadata - a model defined in Swagger"""  # noqa: E501

        self._queuing_budget = None
        self._on_failure = None
        self.discriminator = None

        if queuing_budget is not None:
            self.queuing_budget = queuing_budget
        if on_failure is not None:
            self.on_failure = on_failure

    @property
    def queuing_budget(self):
        """Gets the queuing_budget of this CoreWorkflowMetadata.  # noqa: E501

        Indicates the runtime priority of workflow executions.  # noqa: E501

        :return: The queuing_budget of this CoreWorkflowMetadata.  # noqa: E501
        :rtype: CoreQualityOfService
        """
        return self._queuing_budget

    @queuing_budget.setter
    def queuing_budget(self, queuing_budget):
        """Sets the queuing_budget of this CoreWorkflowMetadata.

        Indicates the runtime priority of workflow executions.  # noqa: E501

        :param queuing_budget: The queuing_budget of this CoreWorkflowMetadata.  # noqa: E501
        :type: CoreQualityOfService
        """

        self._queuing_budget = queuing_budget

    @property
    def on_failure(self):
        """Gets the on_failure of this CoreWorkflowMetadata.  # noqa: E501

        Defines how the system should behave when a failure is detected in the workflow execution.  # noqa: E501

        :return: The on_failure of this CoreWorkflowMetadata.  # noqa: E501
        :rtype: WorkflowMetadataOnFailurePolicy
        """
        return self._on_failure

    @on_failure.setter
    def on_failure(self, on_failure):
        """Sets the on_failure of this CoreWorkflowMetadata.

        Defines how the system should behave when a failure is detected in the workflow execution.  # noqa: E501

        :param on_failure: The on_failure of this CoreWorkflowMetadata.  # noqa: E501
        :type: WorkflowMetadataOnFailurePolicy
        """

        self._on_failure = on_failure

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(CoreWorkflowMetadata, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, CoreWorkflowMetadata):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
