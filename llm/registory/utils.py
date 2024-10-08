from datetime import datetime

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

    def delete(self, collection, query):
        return self.db[collection].delete_one(query)


def record_question(request, message, partial_message):
    # 获取当前时间
    now = datetime.now()
    # 格式化时间
    dt_string = now.strftime("%Y-%m-%d %H:%M:%S")

    db = mongodb("localhost", "llm")
    question_id = db.insert("questions",
                            {"question": message,
                             "answer": partial_message,
                             "like": 0,
                             "create_at": dt_string,
                             "comments": [],
                             "comments_content": [],
                             }
                            ).inserted_id
    referer = request.headers["referer"]
    # 确定用户ID和卡片ID的类型
    card_id = referer.split("cardID=")[-1]
    user_id = referer.split("userID=")[-1].split("&")[0]

    _card_id = ObjectId(card_id)
    _user_id = ObjectId(user_id)

    # 查看新问题与当前卡片记录的最后一个问题是否相同
    card = db.db["cards"].find_one({"_id": _card_id})
    if len(card["questions"]) >= 1:
        last_question_id = card["questions"][-1]
        _last_question_id = ObjectId(last_question_id)
        question = db.db["questions"].find_one({"_id": _last_question_id})
        # print(question["question"])
        if question["question"] == message:
            db.update("cards", {"_id": _card_id}, {"$pop": {"questions": 1}})

            db.update("cards", {"_id": _card_id}, {"$pop": {"content": 1}})
            db.update("cards", {"_id": _card_id}, {"$pop": {"content": 1}})

    # print(question_id, card_id)
    db.update("users", {"_id": _user_id}, {"$push": {"questions": question_id}})
    db.update("cards", {"_id": _card_id}, {"$push": {"questions": question_id, "content": "Q:" + message}})
    db.update("cards", {"_id": _card_id}, {"$push": {"content": "A:" + partial_message}})

