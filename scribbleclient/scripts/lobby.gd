extends Control
@onready var room_id_label: Label = %RoomIdLabel
@onready var start_button: Button = %StartButton

func _ready() -> void:
	room_id_label.text = GlobalState.user._room_id


func _on_copy_button_pressed() -> void:
	DisplayServer.clipboard_set(GlobalState.user._room_id)
	Toast.show_toast("Room Id Copied",0,4)


func _on_start_button_pressed() -> void:
	print("button pressed")
	LoadManager.load_screen("res://scenes/drawing_board.tscn")
