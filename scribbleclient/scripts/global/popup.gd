extends Control

@export var children:PackedScene
@export var min_width:int
@export var min_height:int

signal open
signal close

func _ready() -> void:
	open.connect(_open)
	close.connect(_close)

func _open():
	pass

func _close():
	pass
