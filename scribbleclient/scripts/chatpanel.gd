extends Control
@onready var container: VBoxContainer = $ScrollContainer/VBoxContainer

var message_scene:PackedScene = load("res://scenes/message_item.tscn")

func _ready() -> void:
	GlobalState.new_message_added.connect(add_message)

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
	var scroll = $ScrollContainer
	scroll.scroll_vertical = scroll.get_v_scroll_bar().max_value
