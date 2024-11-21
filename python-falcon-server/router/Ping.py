class Ping:
    @staticmethod
    def on_get(req, res):
        res.text = 'Pong'
