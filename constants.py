import os

from user_data import UserData
from dotenv import load_dotenv

load_dotenv()

CREDENTIALS_FNAME='credentials-%s.json'
OFFERS_HASH_FNAME='hash-%s-%s.json'

USERS = [
  UserData(os.getenv('EMAIL'), os.getenv('CHAT_ID'))
]

TELEGRAM_TOKEN = os.getenv('TELEGRAM_TOKEN')
