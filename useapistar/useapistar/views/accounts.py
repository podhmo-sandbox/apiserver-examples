from .. import store


def accounts():
    return [{"id": d["id"], "name": d["name"]} for d in store.get_accounts()]
