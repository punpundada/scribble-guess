extends Node
class_name Protocol

static func create(
	messaage_type:String, 
	data:Dictionary,
	sender_id:String,
	room_id:String
	)->Dictionary:
	return {
		"messaage_type":messaage_type,
		"meta":{
			"sender_id":sender_id,
			"room_id":room_id,
			"time":Time.get_unix_time_from_system()
		},
		"data":data
	}


static func encode(msg:Dictionary)->String:
	return JSON.stringify(msg)

static func decode(msg:String)->Dictionary:
	var data = JSON.parse_string(msg)
	if data == null:
		return {}
	return data
