import psycopg2
from urllib.parse import urlparse
import os
from dotenv import load_dotenv
import random
import time
import csv

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

now = time.time()
with open('datagen/events_new.csv', newline='') as csvfile:
    reader = csv.reader(csvfile, delimiter=',', quotechar='"')
    with conn.cursor() as cur:
        for row in reader:
            cur.execute("INSERT INTO events_insert (id, source, payload, event_timestamp) VALUES (%s, %s, %s, %s)", (row[0], row[1], row[2], row[3]))
    conn.commit()

end = time.time()
print(f"Total time for insert: {end - now:.2f} seconds")
