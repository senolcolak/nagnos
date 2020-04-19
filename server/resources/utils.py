from flask_restful import Resource


class Ping(Resource):
    def get(self):
        return {'message': 'Live'}, 202

class Howto(Resource):
    def get(self):
        return {'message': 'Live'}, 202
