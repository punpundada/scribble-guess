extends Node

signal messages_changed
signal new_message_added(msg:String)

var messsages:Array[String]=[]

func add_message(msg:String):
	messsages.append(msg)
	messages_changed.emit()
	new_message_added.emit(msg)
