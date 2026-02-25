extends Node


signal progress_changed(progress:float)
signal load_done

var _load_screen_path:String= "res://scenes/global/LoadingScreen.tscn"
var _loading_screen = load(_load_screen_path)
var _loaded_resource:PackedScene
var _scene_path:String
var _progress:Array[float]=[]

var use_sub_threads = true

func load_screen(screen_path:String)->void:
	_scene_path = screen_path
	var new_loading_screen:LoadingScreen = _loading_screen.instantiate()
	self.get_tree().get_root().add_child(new_loading_screen)
	
	self.progress_changed.connect(new_loading_screen._update_progress_bar)
	self.load_done.connect(new_loading_screen._start_outro_animation)
	await Signal(new_loading_screen,"loading_screen_has_full_coverage")
	start_load()

func start_load():
	var state := ResourceLoader.load_threaded_request(_scene_path,"",use_sub_threads)
	if state == OK:
		set_process(true)
	
func _process(_delta: float) -> void:
	var load_status := ResourceLoader.load_threaded_get_status(_scene_path,_progress)
	match load_status:
		0,2:#THREAD LOADED INVALID RESOURCE, THREAD LOAD FAILED
			set_process(false)
			return
		1:self.progress_changed.emit(_progress[0])
		3:
			_loaded_resource = ResourceLoader.load_threaded_get(_scene_path)
			self.progress_changed.emit(1)
			self.load_done.emit()
			self.get_tree().change_scene_to_packed(_loaded_resource)
			set_process(false)
			
			
			
	
	
	
	
	
	
	
	
	
