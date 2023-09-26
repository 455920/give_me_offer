from flask import Flask, render_template, request
import functions

app = Flask(__name__)


@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        input_text = request.form['input_text']
        submit_button = request.form['submit_button']
        output_text = functions.run(input_text, submit_button)
        return render_template('index.html', input_text=input_text, output_text=output_text,
                               buttons=functions.get_buttons())
    else:
        return render_template('index.html', buttons=functions.get_buttons())


if __name__ == '__main__':
    functions.register()
    app.run(debug=True, host="0.0.0.0", port=12345)
