extends Node
class_name MessageRouter

var handlers:Dictionary[String,Callable]={}

func register(type:String, callback:Callable):
	handlers[type]=callback
	
func route(message:Dictionary):
	var t = message.get("message_type","")
	if handlers.has(t):
		handlers[t].call(message)
	else:
		print("Handler does not exis for message_type ",t)
