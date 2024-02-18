from pymongo import MongoClient
from bson.objectid import ObjectId


class mongodb:
    def __init__(self, host, db, port=27017):
        self.host = host
        self.db = db
        self.port = port
        self.client = MongoClient(self.host, self.port)
        self.db = self.client[self.db]

    def insert(self, collection, data):
        return self.db[collection].insert_one(data)

    def update(self, collection, query, data):
        return self.db[collection].update_one(query, data)

    def initialize(self, collection, query):
        if self.db[collection].find_one(query) is None:
            return self.db[collection].update_one(query, {"$set": query}, upsert=True)

def registory_operation(request, message, partial_message):
    db = mongodb("localhost", "llm")
    question_id = db.insert("questions",
                    {"question": message,
                     "answer": partial_message}
                    ).inserted_id
    referer = request.headers["referer"]
    # 确定用户ID和卡片ID的类型
    card_id = referer.split("cardID=")[-1]
    user_id = referer.split("userID=")[-1].split("&")[0]

    _card_id = ObjectId(card_id)
    _user_id = ObjectId(user_id)

    # print(question_id, card_id)
    db.update("users", {"_id": _user_id}, {"$push": {"questions": question_id}})
    db.update("cards", {"_id": _card_id}, {"$push": {"questions": question_id}})
