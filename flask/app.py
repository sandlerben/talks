from flask import Flask, request

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        name = request.form['name']
        return 'Hello {}'.format(name)
    return 'food'

if __name__ == '__main__':
    app.run()
