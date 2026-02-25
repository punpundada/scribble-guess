extends Node
class_name EnvironmentVariable

var API_BASE_URL:String="http://localhost:7421/"

func _init(data:Dictionary) -> void:
	API_BASE_URL = data["API_BASE_URL"]
