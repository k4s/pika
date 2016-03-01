##介绍：
pika是用golang写的一个消息队列，中间件基于redis，RabbitMQ.<br/>

有点类似于python的Celery，但是更加简单，好用。<br/>

提供restful接口异步添加任务。<br/>

##使用：


###building a redis restfulAPI server:
> go run run.go -broker="redis://127.0.0.1:6379/0" -http="127.0.0.1:7778"

###client Demo:

publisher/subscriber 发布订阅通过主题"kas", producer/consumer直接的生产消费通过"me" 
> go run run.go -worker=do -topic="kas" -direct="me" -broker="redis://127.0.0.1:6379/0"

publisher/subscriber 发布订阅通过主题"kas"
> go run run.go -worker=add -topic="kas" -broker="redis://127.0.0.1:6379/0"



###building a rabbitmq restfulAPI server:
> go run run.go -broker="amqp://guest:guest@localhost:5672/" -http="127.0.0.1:7778"


###client Demo:

publisher/subscriber 发布订阅通过主题"kas", producer/consumer直接的生产消费通过"me"
> go run run.go -worker=do -topic="kas" -direct="me" -broker="amqp://guest:guest@localhost:5672/"

publisher/subscriber发布订阅通过主题"kas"
> go run run.go -worker=add -topic="kas" -broker="amqp://guest:guest@localhost:5672/"


###Tasks：

把你的任务写到 github.com/k4s/pika/tasks. <br/>

通过restfulAPI下发任务执行命令：<br/>


> curl -i -X POST http://127.0.0.1:7778/topic/kas -d "tasksdata"

> curl -i -X POST http://127.0.0.1:7778/direct/me -d "tasksdata"