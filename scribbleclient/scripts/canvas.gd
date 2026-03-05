extends Panel
class_name DrawingCanvas

enum tool {PEN,RECTANGLE,CIRCLE,LINE}

signal change_tool(tool)
signal change_line_width(width:float)
signal change_color(color:Color)
signal change_bg_color(color:Color)

#settings
@export var current_color:Color = Color.BLACK
@export var line_width:float = 2.0
@export var drawing:bool = false
@export var current_tool:tool = tool.PEN
@export var canvas_bg_color := Color.WHITE

#store points for freehand drawing
var points := PackedVector2Array()

var start_pos := Vector2.ZERO
var end_pos := Vector2.ZERO

var completed_strokes := []


func _ready() -> void:
	change_tool.connect(set_tool)
	change_line_width.connect(set_line_width)
	change_bg_color.connect(set_background_color)
	change_color.connect(set_color)
	
func _draw() -> void:
	draw_rect(
		Rect2(
			Vector2.ZERO,
			self.get_viewport_rect().size
			),
		canvas_bg_color,true
		)
	
	for stroke in completed_strokes:
		match stroke.type:
			tool.PEN:
				_draw_polyline(stroke.points,stroke.color,stroke.width)
			tool.RECTANGLE:
				draw_rect(stroke.rect,stroke.color,stroke.filled,stroke.width)
			tool.CIRCLE:
				draw_circle(stroke.center, stroke.radius,stroke.color,stroke.filled)
				#if !stroke.filled:
					#var steps = 36
					#var angle_step = 2 *PI / steps
					#for i in steps:
						#var a1 = i+ angle_step
						#var a2 = (i+1) * angle_step
						#var p1 = stroke.center + Vector2(cos(a1),sin(a1)) * stroke.radius
						#var p2 = stroke.center + Vector2(cos(a2),sin(a2)) * stroke.radius
						#draw_line(p1,p2,stroke.color,stroke.width)
			tool.LINE:
				draw_line(stroke.start,stroke.end,stroke.color,stroke.width)
	if drawing and (current_tool in [tool.RECTANGLE,tool.CIRCLE,tool.LINE]):
		match current_tool:
			tool.RECTANGLE:
				var rect =Rect2(start_pos, end_pos-start_pos)
				draw_rect(rect,current_color,false,line_width)
			tool.CIRCLE:
				var radius = start_pos.direction_to(end_pos).length()
				draw_arc(start_pos,radius,0,2*PI,36,current_color,line_width)
			tool.LINE:
				draw_line(start_pos,end_pos,current_color,line_width)
	
	if drawing and current_tool == tool.PEN and points.size() > 1:
		_draw_polyline(points,current_color,line_width)

func _draw_polyline(_points:PackedVector2Array, color:Color,width:float):
	if _points.size() < 2:
		return
	for i in range(_points.size()-1):
		draw_line(_points[i],_points[i+1],color,width,true)


func _input(event: InputEvent) -> void:
	if event is InputEventMouseButton and event.button_index == MOUSE_BUTTON_LEFT:
		var local_pos := self.get_local_mouse_position()
		if event.is_pressed():
			drawing=true
			start_pos=local_pos
			points.clear()
			points.append(local_pos)
			queue_redraw()
		else:
			if drawing:
				match current_tool:
					tool.PEN:
						if points.size() > 1:
							completed_strokes.append({
								"type":tool.PEN,
								"points":points.duplicate(),
								"color":current_color,
								"width":line_width
							})
					tool.RECTANGLE,tool.CIRCLE,tool.LINE:
						end_pos = self.get_local_mouse_position()
						if current_tool == tool.RECTANGLE:
							completed_strokes.append({
								"type":tool.RECTANGLE,
								"rect":Rect2(start_pos,end_pos-start_pos),
								"color":current_color,
								"filled":false,
								"width":line_width
							})
						if current_tool == tool.CIRCLE:
							completed_strokes.append({
								"type":tool.CIRCLE,
								"center":start_pos,
								"radius": start_pos.distance_to(end_pos),
								"color":current_color,
								"filled":false,
								"width":line_width
							})
						if current_tool == tool.LINE:
							completed_strokes.append({
								"type":tool.LINE,
								"start":start_pos,
								"end":end_pos,
								"color":current_color,
								"width":line_width
							})
			drawing = false
			queue_redraw()
	elif event is InputEventMouseMotion and drawing:
		var local_pos := self.get_local_mouse_position()
		if current_tool == tool.PEN:
			points.append(local_pos)
		else:
			end_pos = local_pos
		queue_redraw()


func set_tool(tool_name:tool):
	current_tool=tool_name

func set_color(color:Color):
	current_color=color

func set_line_width(width:float):
	line_width=width
	
func set_background_color(color:Color):
	canvas_bg_color=color
	queue_redraw()

func clear_canvas():
	completed_strokes.clear()
	points.clear()
	drawing=false
	queue_redraw()
