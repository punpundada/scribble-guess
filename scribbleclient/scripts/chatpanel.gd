extends Control
@onready var container: VBoxContainer = $VBoxContainer/ScrollContainer/VBoxContainer
@onready var message_text_box: TextEdit = $VBoxContainer/MarginContainer/VBoxContainer/MessageTextBox
@onready var drawing_canvas: DrawingCanvas = $"../../DrawingCanvas"

var message_scene:PackedScene = load("res://scenes/message_item.tscn")
var menu_scene:PackedScene = load("res://scenes/menu.tscn")
var text_msg:String

func _ready() -> void:
	GlobalState.new_message_added.connect(add_message)
	message_text_box.text_changed.connect(on_text_changed)
	WebsocketManager.disconnect.connect(_change_scene_to_menu)

func render_messages(messages:Array):
	clear_messages()
	for msg in messages:
		add_message(msg)

func add_message(msg:String):
	print("msg12 ",msg)
	var item:MessageItem = message_scene.instantiate()
	item.set_message(msg)
	container.add_child(item)

	await get_tree().process_frame
	scroll_to_bottom()
	
func clear_messages():
	for c in container.get_children():
		c.queue_free()
		
func scroll_to_bottom():
	var scroll = $VBoxContainer/ScrollContainer
	scroll.scroll_vertical = scroll.get_v_scroll_bar().max_value

func on_text_changed():
	text_msg=message_text_box.text

func _on_send_button_pressed() -> void:
	if len(text_msg)==0:
		return
	WebsocketManager.send(
		"message",
		text_msg,
		GlobalState.user._id,
		GlobalState.user._room_id
		)
	text_msg=""
	message_text_box.text=""

func _change_scene_to_menu():
	Toast.show_toast("Websocket Disconnected",10)
	print("websocket disconnected")
	
