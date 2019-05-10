package tech.thomaslegrand.h2o

import io.ktor.application.Application
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import org.slf4j.Logger
import org.slf4j.LoggerFactory

val LOG: Logger = LoggerFactory.getLogger("ktor-server")

fun main(args: Array<String>) {

    LOG.debug("Starting ktor server.")

    embeddedServer(Netty, 8081, module = Application::main).start(wait = true)


}