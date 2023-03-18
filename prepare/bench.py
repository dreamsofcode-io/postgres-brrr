import psycopg2
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

conn = psycopg2.connect(**pg_connection_dict)


MAX_LIMIT = 255
MAX_NAME_LEN = 80

fake = Faker()
data = []

for i in range(100000):
    name = ''

    len = random.randint(0, MAX_NAME_LEN)
    for i in range(len):
        random_integer = random.randint(0, MAX_LIMIT)
        name += (chr(random_integer))

    name = fake.job()
    date = fake.date_between(start_date='today', end_date='+30y')
    is_done = bool(random.getrandbits(1))

    data.append((name, date, is_done))

cur = conn.cursor()
cur.execute("""
PREPARE insert_task AS 
INSERT INTO task (name, due_date, is_done)
VALUES ($1, $2, $3)
""")
cur.close()

start = time.time()
for row in data:
    cur = conn.cursor()
    cur.execute("EXECUTE insert_task (%s, %s, %s)", row, prepare=False)
    cur.close()

end = time.time()
print(f"using prepared statements: {end - start : .2f} seconds")

start = time.time()
for row in data:
    cur = conn.cursor()
    cur.execute("INSERT INTO task (name, due_date, is_done) VALUES (%s, %s, %s)", row, prepare=False)
    cur.close()

end = time.time()
print(f"direct queries: {end - start : .2f} seconds")

cur = conn.cursor()
cur.execute("DELETE FROM task")

conn.commit()
