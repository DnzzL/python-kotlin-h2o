package tech.thomaslegrand.h2o


import hex.genmodel.easy.EasyPredictModelWrapper
import hex.genmodel.easy.RowData
import hex.genmodel.easy.prediction.MultinomialModelPrediction
import hex.genmodel.ModelMojoReader
import hex.genmodel.MojoReaderBackendFactory


class Model {
    fun convert_to_raw_data(request: Request): RowData {
        val test = hashMapOf(
            "SepalLength" to request.SepalLength,
            "SepalWidth" to request.SepalWidth,
            "PetalLength" to request.PetalLength,
            "PetalWidth" to request.PetalWidth
        )
        val testRow = RowData()
        test.forEach { (k, v) ->
            testRow[k] = v
        }
        return testRow
    }

    fun predict(request: Request): String? {
        val mojoURL = this::class.java.classLoader.getResource("rf_fit.zip")
        val reader =
            MojoReaderBackendFactory.createReaderBackend(mojoURL, MojoReaderBackendFactory.CachingStrategy.MEMORY)
        val model = ModelMojoReader.readFrom(reader)
        val modelWrapper = EasyPredictModelWrapper(model)

        val testRow = convert_to_raw_data(request)
        val prediction = modelWrapper.predict(testRow) as MultinomialModelPrediction
        return prediction.label
    }
}

fun main() {
    val test = Request(2.0, 1.0, 1.5, 2.4)
    val prediction = Model().predict(test)
    println(prediction)
}