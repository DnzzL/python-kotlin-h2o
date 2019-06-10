package tech.thomaslegrand.h2o


import hex.genmodel.easy.EasyPredictModelWrapper
import hex.genmodel.easy.RowData
import hex.genmodel.easy.prediction.MultinomialModelPrediction
import hex.genmodel.ModelMojoReader
import hex.genmodel.MojoReaderBackendFactory


class Model {
    private var model: EasyPredictModelWrapper

    init {
        val mojoURL = this::class.java.classLoader.getResource("rf_fit.zip")
        val reader =
            MojoReaderBackendFactory.createReaderBackend(mojoURL, MojoReaderBackendFactory.CachingStrategy.MEMORY)
        val mojo = ModelMojoReader.readFrom(reader)
        model = EasyPredictModelWrapper(mojo)
    }


    private fun convertToRawData(request: IrisRequest): RowData {
        val test = hashMapOf(
            "SepalLength" to request.sepalLength,
            "SepalWidth" to request.sepalWidth,
            "PetalLength" to request.petalLength,
            "PetalWidth" to request.petalWidth
        )
        val testRow = RowData()
        test.forEach { (k, v) ->
            testRow[k] = v
        }
        return testRow
    }


    fun predict(request: IrisRequest): String? {
        val testRow = convertToRawData(request)
        val prediction = model.predict(testRow) as MultinomialModelPrediction
        print(prediction.label)
        return prediction.label
    }
}