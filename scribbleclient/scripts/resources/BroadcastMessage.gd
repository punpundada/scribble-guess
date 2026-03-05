extends Resource
class_name BroadcastMessage

var _messaage_type:String
var _data:Variant
var _meta:Meta

func _init(message_type:String,meta:Meta,data:Variant) -> void:
	_messaage_type=message_type
	_meta=meta
	_data=data
