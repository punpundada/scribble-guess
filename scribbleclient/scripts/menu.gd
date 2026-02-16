extends Control

var texture = preload("res://assets/PNG/cursorGauntlet_blue.png")

func _ready() -> void:
	Input.set_custom_mouse_cursor(texture)
