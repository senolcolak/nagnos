from flask_restful import Resource, reqparse
from flask_jwt import jwt_required
from models.command import CommandModel


class Command(Resource):
    parser = reqparse.RequestParser()
    parser.add_argument('explanation',
                        type=str,
                        required=True,
                        help="This field cannot be left blank!"
                        )
    parser.add_argument('category_id',
                        type=int,
                        required=True,
                        help="Every item needs a category_id."
                        )

    #@jwt_required()
    def get(self, name):
        command = CommandModel.find_by_name(name)
        if command:
            return command.json()
        return {'message': 'Command not found'}, 404

    def post(self, name):
        if CommandModel.find_by_name(name):
            return {'message': "A command with name '{}' already exists.".format(name)}, 400

        data = command.parser.parse_args()

        command = CommandModel(name, **data)

        try:
            command.save_to_db()
        except:
            return {"message": "An error occurred inserting the command."}, 500

        return command.json(), 201

    def delete(self, name):
        command = CommandModel.find_by_name(name)
        if command:
            command.delete_from_db()
            return {'message': 'Command deleted.'}
        return {'message': 'Command not found.'}, 404

    def put(self, name):
        data = Command.parser.parse_args()

        command = CommandModel.find_by_name(name)

        if command:
            command.explanation = data['explanation']
        else:
            command = CommandModel(name, **data)

        command.save_to_db()

        return command.json()


class CommandList(Resource):
    def get(self):
        return {'commands': list(map(lambda x: x.json(), CommandModel.query.all()))}

class CommandListShort(Resource):
    def get(self):
        return {'commands': list(map(lambda x: x, CommandModel.query.with_entities(CommandModel.name)))}
