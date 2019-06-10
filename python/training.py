import h2o
from h2o.estimators import H2ORandomForestEstimator

h2o.init(nthreads=-1, max_mem_size=16)

iris_path = "data/iris.csv"
data = h2o.import_file(iris_path)
print(f'data shape -- {data.shape}')

levels = data['Species'].levels()
print(f'encoded classes -- {levels}')

splits = data.split_frame(ratios=[0.6, 0.2], seed=1)

train = splits[0]
valid = splits[1]
test = splits[2]

y = 'Species'
x = list(data.columns)
x.remove(y)

rf_fit = H2ORandomForestEstimator(model_id='rf_fit', seed=1, balance_classes=True, nfolds=3)
rf_fit.train(x=x, y=y, training_frame=train, validation_frame=valid)
rf_perf = rf_fit.model_performance(test)
print(rf_perf)

rf_fit.save_mojo(f"models")
h2o.save_model(rf_fit, f"models")
