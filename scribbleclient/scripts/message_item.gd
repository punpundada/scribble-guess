extends Panel
class_name MessageItem

func set_message(message:String):
	self.get_children().get(0).text = message
