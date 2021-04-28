import psycopg2
import boto3
import random

conn = psycopg2.connect(
    host="prod-1.csz2smpfztf7.us-east-1.rds.amazonaws.com",
    database="maindb",
    user="registerio",
    password="registera")

cur = conn.cursor()

cur.execute("SELECT course_registrations.netid, course_registrations.class_index,\"sqs queues\".url FROM course_registrations, \"sqs queues\" WHERE course_registrations.class_index = \"sqs queues\".index;")

rows = cur.fetchall()

client = boto3.client('sqs')

for row in rows:
    if "bot" in row[0]:
        response = client.send_message(
            QueueUrl=row[2],
            MessageBody="{}|drop".format(row[0]),
            DelaySeconds=0,
            MessageDeduplicationId=row[0]+row[2]+"{}".format(random.randrange(10000)),
            MessageGroupId='delete'
        )