extends Control

@onready var http_request: HTTPRequest = $HTTPRequest
@onready var play_button: Button = $VBoxContainer/HBoxContainer/PlayButton

var texture = preload("res://assets/PNG/cursorGauntlet_blue.png")

func _ready() -> void:
	Input.set_custom_mouse_cursor(texture)
	WebsocketManager.router.register(
		"message",
		func(msg):GlobalState.add_message(msg["data"])
	)
	

func _on_play_button_pressed() -> void:
	play_button.disabled = true
	http_request.request(
		"http://localhost:7421/api/room/random",
	)
	http_request.request_completed.connect(_on_request_complete)
	
func _on_request_complete(_result,_response_code,_headers,body):
	var json = JSON.parse_string(body.get_string_from_utf8())
	print("roomId ",json["roomId"])
	WebsocketManager.connect_socket(json["roomId"])
	LoadManager.load_screen("res://scenes/drawing_board.tscn")
	play_button.disabled = false
	
