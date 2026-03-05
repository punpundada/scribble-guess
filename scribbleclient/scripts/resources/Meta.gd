extends Resource
class_name Meta

var _sender_id:String
var _room_id:String
var _time:int

func _init(sender_id:String,room_id:String,time:int) -> void:
	_sender_id=sender_id
	_room_id = room_id
	_time = time
	
