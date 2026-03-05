extends Node

signal messages_changed
signal new_message_added(msg:String)
signal user_connected(user:Dictionary)

var messsages:Array[String]=[]
var user:User = null

func add_message(msg:String):
	messsages.append(msg)
	messages_changed.emit()
	new_message_added.emit(msg)

func user_connected_romm(connected_user:Dictionary):
	user_connected.emit(connected_user)
	user = User.new(
		connected_user.get("name","Default name"),
		connected_user["meta"]["sender_id"],
		connected_user["meta"]["room_id"]
		)
	
