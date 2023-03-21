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
    with conn.cursor() as cur:
        cur.copy_from(csvfile, 'events_copy', sep=',', columns=('id', 'source', 'payload', 'event_timestamp'))
        conn.commit()

end = time.time()

conn.close()

print(f"Total time for copy: {end - now:.2f} seconds")
