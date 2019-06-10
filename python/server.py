import time
from concurrent import futures

import fire
import grpc
import numpy as np
from h2o import h2o, H2OFrame

import Iris_pb2_grpc
from Iris_pb2 import IrisReply, IrisRequest

_ONE_DAY_IN_SECONDS = 60 * 60 * 24
SPECIES = {
    1: "Iris-setosa",
    2: "Iris-setosa",
    3: "Iris-setosa"
}


class PredictorServicer(Iris_pb2_grpc.PredictorServicer):

    def __init__(self, model_path):
        self.model = h2o.load_model(model_path)

    def Predict(self, request: IrisRequest, context):
        if not hasattr(request, 'SepalLength') or not hasattr(request, 'SepalWidth') \
                or not hasattr(request, 'PetalLength') or not hasattr(request, 'PetalWidth'):
            msg = 'wrong arguments for IrisRequest'
            context.set_details(msg)
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)

        test_data = H2OFrame({"SepalLength": request.SepalLength, "SepalWidth": request.SepalWidth,
                              "PetalLength": request.PetalLength, "PetalWidth": request.PetalWidth})
        prediction = self.model.predict(test_data).getrow()
        species = SPECIES.get(np.argmax(prediction))
        return IrisReply(species=species)


def serve(model_path):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    Iris_pb2_grpc.add_PredictorServicer_to_server(
        PredictorServicer(model_path), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


def main(model_path):
    h2o.init(nthreads=-1, max_mem_size=16)
    serve(model_path)


if __name__ == '__main__':
    fire.Fire(main)
