package tech.thomaslegrand.h2o

import io.grpc.ManagedChannel
import io.grpc.ManagedChannelBuilder
import io.grpc.StatusRuntimeException
import java.util.concurrent.TimeUnit
import java.util.logging.Level
import java.util.logging.Logger

/**
 * A simple client that requests a greeting from the [IrisServer].
 */
class IrisClient
/** Construct client for accessing RouteGuide server using the existing channel.  */
internal constructor(private val channel: ManagedChannel) {
    private val blockingStub: PredictorGrpc.PredictorBlockingStub = PredictorGrpc.newBlockingStub(channel)

    /** Construct client connecting to server at `host:port`.  */
    constructor(host: String, port: Int) : this(
        ManagedChannelBuilder.forAddress(host, port)
            // Channels are secure by default (via SSL/TLS). For the example we disable TLS to avoid
            // needing certificates.
            .usePlaintext()
            .build()
    ) {}


    @Throws(InterruptedException::class)
    fun shutdown() {
        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS)
    }

    /** Send request to server.  */
    fun getPrediction(sepalLength: Double, sepalWidth: Double, petalLength: Double, petalWidth: Double) {
        logger.log(Level.INFO, "Will try to get predictions")
        val request = IrisRequest.newBuilder()
            .setPetalLength(petalLength)
            .setPetalWidth(petalWidth)
            .setSepalLength(sepalLength)
            .setSepalWidth(sepalWidth).build()
        val response: IrisReply = try {
            blockingStub.predict(request)
        } catch (e: StatusRuntimeException) {
            logger.log(Level.WARNING, "RPC failed: {0}", e.status)
            return
        }

        logger.info("Predicted: ${response.species}")
    }

    companion object {
        private val logger = Logger.getLogger(IrisClient::class.java.name)

        /**
         * Server. If provided, the first four elements of `args` are the
         * measure of the flower.
         */
        @Throws(Exception::class)
        @JvmStatic
        fun main(args: Array<String>) {
            val client = IrisClient("localhost", 50051)
            try {
                /* Access a service running on the local machine on port 50051 */
                val sepalLength = if (args.isNotEmpty()) args[0].toDouble() else 1.1
                val sepalWidth= if (args.isNotEmpty()) args[1].toDouble() else 1.4
                val petalLength= if (args.isNotEmpty()) args[2].toDouble() else 1.3
                val petalWidth= if (args.isNotEmpty()) args[3].toDouble() else 1.8
                client.getPrediction(sepalLength,sepalWidth,petalLength,petalWidth)
            } finally {
                client.shutdown()
            }
        }
    }
}

