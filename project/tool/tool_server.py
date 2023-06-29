from flask import Flask, render_template, request
import base64

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        input_text = request.form['input_text']
        output_text = request.form['output_text']
        if request.form['submit_button'] == 'Encode':
            output_text = base64.b64encode(input_text.encode('utf-8')).decode('utf-8')
        elif request.form['submit_button'] == 'Decode':
            output_text = base64.b64decode(input_text.encode('utf-8')).decode('utf-8')
        return render_template('index.html', input_text=input_text, output_text=output_text)
    else:
        return render_template('index.html')

if __name__ == '__main__':
    app.run(debug=True)