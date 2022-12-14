# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import python_service_list_pb2 as python__service__list__pb2


class EventsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListOneItem = channel.unary_unary(
                '/Events/ListOneItem',
                request_serializer=python__service__list__pb2.Item.SerializeToString,
                response_deserializer=python__service__list__pb2.ResponseValue.FromString,
                )


class EventsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ListOneItem(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_EventsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListOneItem': grpc.unary_unary_rpc_method_handler(
                    servicer.ListOneItem,
                    request_deserializer=python__service__list__pb2.Item.FromString,
                    response_serializer=python__service__list__pb2.ResponseValue.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Events', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Events(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ListOneItem(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Events/ListOneItem',
            python__service__list__pb2.Item.SerializeToString,
            python__service__list__pb2.ResponseValue.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
