import psycopg
from urllib.parse import urlparse
import os
from faker import Faker
from dotenv import load_dotenv
import random
import time

load_dotenv()

conStr = os.getenv('DATABASE_URL', 'postgres://postgres:postgres@localhost:5432/postgres')
p = urlparse(conStr)

pg_connection_dict = {
    'dbname': p.path[1:],
    'user': p.username,
    'password': p.password,
    'port': p.port,
    'host': p.hostname
}

conn = psycopg.connect(**pg_connection_dict)

for i in range(2000000):
    cur = conn.cursor()
    cur.execute("INSERT INTO users (name, age) VALUES (%s, %s)", (Faker().name(), random.randint(0, 100)))
    conn.commit()
    cur.close()
