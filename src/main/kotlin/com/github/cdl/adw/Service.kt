package com.github.cdl.adw

import it.uniroma1.lcl.adw.*
import it.uniroma1.lcl.adw.ADW
import it.uniroma1.lcl.adw.ItemType
import it.uniroma1.lcl.adw.utils.*
import it.uniroma1.lcl.adw.comparison.*
import io.javalin.apibuilder.ApiBuilder.*
import io.javalin.Javalin
import io.javalin.Context

import java.io.File


fun main(args: Array<String>) {
    val localDir = File(".").absolutePath
    System.out.println("running at path: $localDir")

    val app = Javalin.create().start(8080)

    app.routes {
        get("/healthz") { ctx -> ctx.result("{\"status\": \"ok\"}") }
        path("/argument/adw") {
            post(ADWHandler::handle)
        }
    }
}

object ADWHandler {

    val srcTextType = ItemType.SURFACE
    val trgTextType = ItemType.SURFACE
    val disMethod = DisambiguationMethod.ALIGNMENT_BASED
    val measure = WeightedOverlap()
    val pipeLine = ADW.getInstance()

    private data class Request(val text1: String = "", val text2: String = "")
    private data class Response(val value: Double = 0.0)

    fun handle(ctx: Context) {
        val req = ctx.bodyAsClass(Request::class.java)
        val v = pipeLine.getPairSimilarity(
          req.text1, req.text2,
          disMethod, measure,
          srcTextType, trgTextType)
        
        getResponse(ctx, v)
    }

    fun getResponse(ctx: Context, v: Double) {
        ctx.json(Response(v))
    }

}
