import base64
import hashlib

button_to_function = {}

def base64_encode(input_text):
    output_text = base64.b64encode(input_text.encode('utf-8')).decode('utf-8')
    return output_text

def base64_decode(input_text):
    output_text = base64.b64decode(input_text.encode('utf-8')).decode('utf-8')
    return output_text

def calc_md5(input_text):
    m = hashlib.md5()
    m.update(input_text.encode('utf-8'))
    output_text = m.hexdigest()
    return output_text

def add_button(button, function):
    button_to_function[button] = function

def register():
    add_button("Encode", base64_encode)
    add_button("Decode", base64_decode)
    add_button("MD5", calc_md5)

def get_buttons():
    return button_to_function


def run(input_text, submit_button):
    try:
        return button_to_function[submit_button](input_text)
    except Exception as e:
        return str(e)