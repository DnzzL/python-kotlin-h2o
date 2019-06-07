# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import Iris_pb2 as Iris__pb2


class PredictorStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Predict = channel.unary_unary(
        '/h2o.Predictor/Predict',
        request_serializer=Iris__pb2.IrisRequest.SerializeToString,
        response_deserializer=Iris__pb2.IrisReply.FromString,
        )


class PredictorServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Predict(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_PredictorServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Predict': grpc.unary_unary_rpc_method_handler(
          servicer.Predict,
          request_deserializer=Iris__pb2.IrisRequest.FromString,
          response_serializer=Iris__pb2.IrisReply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'h2o.Predictor', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
