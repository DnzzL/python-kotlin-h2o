package tech.thomaslegrand.h2o

import io.grpc.Server
import io.grpc.ServerBuilder
import io.grpc.stub.StreamObserver
import java.io.IOException
import java.util.logging.Level
import java.util.logging.Logger

class IrisServer {

    private var server: Server? = null

    @Throws(IOException::class)
    private fun start() {
        /* The port on which the server should run */
        val port = 50051
        server = ServerBuilder.forPort(port)
            .addService(PredictorImpl())
            .build()
            .start()
        logger.log(Level.INFO, "Server started, listening on {0}", port)
        Runtime.getRuntime().addShutdownHook(object : Thread() {
            override fun run() {
                // Use stderr here since the logger may have been reset by its JVM shutdown hook.
                System.err.println("*** shutting down gRPC server since JVM is shutting down")
                this@IrisServer.stop()
                System.err.println("*** server shut down")
            }
        })
    }

    private fun stop() {
        server?.shutdown()
    }

    /**
     * Await termination on the main thread since the grpc library uses daemon threads.
     */
    @Throws(InterruptedException::class)
    private fun blockUntilShutdown() {
        server?.awaitTermination()
    }

    internal class PredictorImpl : PredictorGrpc.PredictorImplBase() {

        private var model: Model = Model()

        override fun predict(request: IrisRequest?, responseObserver: StreamObserver<IrisReply>?) {
            val prediction = model.predict(request!!)
            val reply = IrisReply.newBuilder().setSpecies(prediction).build()
            responseObserver?.onNext(reply)
            responseObserver?.onCompleted()
        }
    }

    companion object {
        private val logger = Logger.getLogger(IrisServer::class.java.name)

        /**
         * Main launches the server from the command line.
         */
        @Throws(IOException::class, InterruptedException::class)
        @JvmStatic
        fun main(args: Array<String>) {
            val server = IrisServer()
            server.start()
            server.blockUntilShutdown()
        }
    }
}