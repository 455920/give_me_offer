import base64
import hashlib
import random

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


def gen_random_temperature(input_text: str):
    if not input_text.isdigit():
        return "输入错误: 必须是一个数字n, 会生成n个随机温度"
    n = int(input_text)
    if n > 128:
        n = 128
    output_text = ''
    for i in range(1, n+1):
        r1 = random.randint(362, 370)
        r2 = random.randint(362, 370)
        output_text += "(%d) %2.1f %2.1f\n" % (i, r1/10.0, r2/10.0)
    return output_text

def gen_table1_sql(input_text: str):
    if not input_text.isdigit():
        return "输入错误: 必须是一个数字"
    uid = int(input_text)
    tb = uid%1000/100
    db = uid%100
    return "db_%02d.table1_%d" % (db, tb)

def gen_table2_sql(input_text: str):
    if not input_text.isdigit():
        return "输入错误: 必须是一个数字"
    uid = int(input_text)
    tb = uid%1000/100
    db = uid%100
    return "db_%02d.table2_%d" % (db, tb)

def add_button(button_name, function):
    button_to_function[button_name] = function


# 增加一个工具按钮,只需要实现在register注册一个按钮, 不需要修改html文件
def register():
    add_button("Base64Encode", base64_encode)
    add_button("Base64Decode", base64_decode)
    add_button("MD5", calc_md5)
    add_button("生成随机正常的人体温度", gen_random_temperature)
    add_button("table1", gen_table1_sql)
    add_button("table2", gen_table2_sql)

def get_buttons():
    return button_to_function


def run(input_text, submit_button):
    try:
        return button_to_function[submit_button](input_text)
    except Exception as e:
        return str(e)
