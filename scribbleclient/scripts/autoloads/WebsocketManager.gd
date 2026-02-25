extends Node
#class_name WebsocketManager

signal connect
signal disconnect
signal message(msg)

@export var URL:String = "ws://localhost:7421/api/ws"

var socket := WebSocketPeer.new()

var state := ConnectionState.State.DISCONNECTED

var router:MessageRouter = MessageRouter.new()

func connect_socket(room_id:String=""):
	if state == ConnectionState.State.CONNECTED:
		return
	state = ConnectionState.State.CONNECTING
	var url := URL
	if len(room_id) > 1:
		url += "?roomId="+room_id
	var err := socket.connect_to_url(url)
	if err != OK:
		print("Unable to connect to WS %s",URL)
	else:
		print("Websocket Connected")
	
func send(type:String, data:Dictionary,sender_id:String,room_id:String):
	if socket.get_ready_state() != WebSocketPeer.STATE_OPEN:
		return
	var msg:Dictionary = Protocol.create(type,data,sender_id,room_id)
	socket.send_text(Protocol.encode(msg))

func _process(_delta: float) -> void:
	socket.poll()
	match socket.get_ready_state():
		WebSocketPeer.STATE_OPEN:
			if state != ConnectionState.State.CONNECTED:
				state = ConnectionState.State.CONNECTED
				connect.emit()
			while socket.get_available_packet_count():
				var text = socket.get_packet().get_string_from_utf8()
				var msg:Dictionary = Protocol.decode(text)
				router.route(msg)
				message.emit(msg)
		WebSocketPeer.STATE_CLOSED:
			if state != ConnectionState.State.DISCONNECTED:
				state = ConnectionState.State.DISCONNECTED
				disconnect.emit()
				set_process(false)
	
func close():
	socket.close()
	
	
	
	
	
	
