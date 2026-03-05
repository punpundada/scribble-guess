extends CanvasLayer

var toast_scene_packed:PackedScene = load("uid://cg82d4n4n4d1r")
var timer:Timer = Timer.new()

func close_toast():
	pass

func show_toast(text:String,delay:int=0,time_sec:int=2):
	if delay >0:
		var delayTimer = Timer.new()
		delayTimer.wait_time=delay
		self.get_tree().get_root().add_child(delayTimer)
		delayTimer.start()
		await delayTimer.timeout
	var toast_scene := toast_scene_packed.instantiate()
	var animation_player = toast_scene.get_node("%ToastAnimationPlayer")
	toast_scene.add_to_group("ToastGroup")
	toast_scene.add_child(timer)
	toast_scene.name = "toast_container"
	self.get_tree().get_root().add_child(toast_scene)
	toast_scene.get_node("%ToastLabel").text=text
	timer.wait_time=time_sec
	timer.start()
	await timer.timeout
	animation_player.play("exit")
	animation_player.animation_finished.connect(
		func(_n):toast_scene.queue_free()
		)
	
	


func _on_close_button_pressed() -> void:
	print("Button pressed")
	if timer != null:
		timer.stop()
		timer.timeout.emit()
	print("stopped")
