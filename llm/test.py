

referer = "http://localhost:8082/?userID=65cf55a99a47e8b0143139b0&cardID=65d0dc61bb41a9345fa71ede"
card_id = referer.split("cardID=")[-1]
user_id = referer.split("userID=")[-1].split("&")[0]

print(card_id)
print(user_id)