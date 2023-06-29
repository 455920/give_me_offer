import base64

button_to_function = {}

def base64_encode(input_text):
    output_text = base64.b64encode(input_text.encode('utf-8')).decode('utf-8')
    return output_text

def base64_decode(input_text):
    output_text = base64.b64decode(input_text.encode('utf-8')).decode('utf-8')
    return output_text

def register():
    button_to_function["Encode"] = base64_encode
    button_to_function["Decode"] = base64_decode

def run(input_text, submit_button):
    print(input_text)
    print(submit_button)
    print(button_to_function[submit_button])
    return button_to_function[submit_button](input_text)
