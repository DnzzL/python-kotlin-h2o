package tech.thomaslegrand.h2o

import io.ktor.application.*
import io.ktor.http.*
import io.ktor.request.receive
import io.ktor.response.*
import io.ktor.routing.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*

const val SENTIMENT_ENDPOINT  = "/sentiment"

data class Request(val SepalLength: Double, val SepalWidth: Double, val PetalLength: Double, val PetalWidth: Double)

fun Application.main() {
        routing {
            get("/") {
                call.respondText("Hello World!")
            }
            post("/sentiment") {
                call.respondText("Hello, world!", ContentType.Text.Html)
//                val request = call.receive<Request>()
//                val test = Request(2.0, 1.0, 1.5, 2.4)
//                val prediction = Model().predict(test)
//                call.respondText(prediction!!)
            }

        }
}
