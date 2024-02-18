from flask import Flask, request

app = Flask(__name__)


@app.route('/')
def hello_world():
    ids = str(request.args.get('?'))
    return 'Hello, World!' + ids
