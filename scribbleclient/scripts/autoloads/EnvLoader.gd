extends Node
var ENV:EnvironmentVariable

func _ready() -> void:
	#load_env("../../../.env")
	pass

func load_env(path:String)->EnvironmentVariable:
	print(get_script().resource_path)
	if not FileAccess.file_exists(path):
		push_error("Env file not found "+path)
		return null
	var file = FileAccess.open(path,FileAccess.READ)
	var data:Dictionary
	while not file.eof_reached():
		var line = file.get_line().strip_edges()
		if line == "" or line.begins_with("#"):
			continue
		var parts = line.split("=",false,1)
		if parts.size() == 2:
			data[parts[0]] = parts[1]
	var env := EnvironmentVariable.new(data)
	return env

func list_files(path:String):
	var dir = DirAccess.open(path)
	if dir == null:
		print("failed to open dir "+path)
		return
	dir.list_dir_begin()
	var file_name = dir.get_next()
	while file_name != "":
		if not dir.current_is_dir():
			print(file_name)
	file_name = dir.get_next()
	dir.list_dir_end()
	
