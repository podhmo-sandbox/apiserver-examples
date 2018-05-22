from apistar import App, Route
from importlib import import_module


def make_app():
    routes = [
        Route('/', method='GET', handler=import_module("useapistar.views.welcome").welcome),
        Route('/cellar/accounts', method='GET', handler=import_module("useapistar.views.accounts").accounts),
    ]
    return App(routes=routes)
