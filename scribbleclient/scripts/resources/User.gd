extends Resource
class_name User

var _name:String=""
var _id:String = ""
var _room_id:String=""

func _init(name:String,id:String,room_id:String) -> void:
	_name = name
	_id = id
	_room_id = room_id
