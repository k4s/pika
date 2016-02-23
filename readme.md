##building a restfulAPI server:
go run run.go -broker="redis://127.0.0.1:6379/0" -http="127.0.0.1:7778

##client Demo:
###publisher/subscriber by "kas", producer/consumer by "me" 
go run run.go -worker=do -topic="kas" -plant="me"  -broker="redis://127.0.0.1:6379/0
###publisher/subscriber by "kas"
go run run.go -worker=add -topic="kas"  -broker="redis://127.0.0.1:6379/0
