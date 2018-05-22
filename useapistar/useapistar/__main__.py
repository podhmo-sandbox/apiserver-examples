from useapistar import make_app
app = make_app()
app.serve('127.0.0.1', 8081, debug=True)
